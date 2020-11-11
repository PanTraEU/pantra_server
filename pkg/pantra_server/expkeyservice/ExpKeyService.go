package expkeyservice

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	configUtil "github.com/pantraeu/pantra_server/pkg/pantra_server/confutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/model/authtoken"
	expkey "github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const BYTES_IN_INT32 = 4

var (
	config = configUtil.GetConfig()
)

func unsafeCaseInt32ToBytes(val int32) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: BYTES_IN_INT32, Cap: BYTES_IN_INT32}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func GetExpKeysByOffset(c *fiber.Ctx) error {
	offset, err := strconv.Atoi(c.Params("offset", "0"))

	if err != nil {
		return err
	}

	page, err := strconv.Atoi(c.Params("page", "0"))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	size, err := strconv.Atoi(c.Params("size", "10"))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	expKeys, err := expkey.GetExpKeysByOffset(offset, page, size)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if len(expKeys) == 0 {
		err = c.SendStatus(404)
		if err != nil {
			log.Error(err.Error())
		}
	} else {
		restKeys := []expkey.ExpKeyRest{}
		for _, eKey := range expKeys {
			rKey := new(expkey.ExpKeyRest)
			rKey.Day = eKey.Day
			rKey.ExpKey = eKey.ExpKey
			restKeys = append(restKeys, *rKey)
		}
		err := c.SendString(createCSV(restKeys, true))
		if err != nil {
			log.Error(err.Error())
		}
	}
	return nil
}

func GetExpKeysByDate(c *fiber.Ctx, bindata bool) error {

	dateStr := c.Params("date")
	if len(dateStr) != 8 {
		return fmt.Errorf("invalid date (yyyymmdd): %s", dateStr)
	}

	page, err := strconv.Atoi(c.Params("page", "0"))
	if err != nil {
		log.Errorf("GetAllExpKeysByDate: %s", err.Error())
		return err
	}

	size, err := strconv.Atoi(c.Params("size", "10"))
	if err != nil {
		log.Errorf("GetAllExpKeysByDate: %s", err.Error())
		return err
	}

	log.Debugf("GetAllExpKeysByDate: %s / % d / %d", dateStr, page, size)

	expKeys, err := expkey.GetAllExpKeysByDate(dateStr, page, size)
	if err != nil {
		log.Errorf("GetAllExpKeysByDate: %s", err.Error())
		return err
	}

	if len(expKeys) == 0 {
		log.Error("GetAllExpKeysByDate: no more keys available")

		if page == 0 { // case no keys at all
			//err := c.SendStatus(http.StatusNotFound)
			err := c.SendStatus(http.StatusNoContent)
			log.Errorf("GetAllExpKeysByDate (no keys): %s", err.Error())
			return err
		} else { // case no more keys
			err := c.SendStatus(http.StatusNotFound)
			log.Errorf("GetAllExpKeysByDate (no more keys): %s", err.Error())
			return err
		}
	} else {
		restKeys := []expkey.ExpKeyRest{}
		for _, eKey := range expKeys {
			rKey := new(expkey.ExpKeyRest)
			rKey.Day = eKey.Day
			rKey.ExpKey = eKey.ExpKey
			rKey.RollingStartIntervalNumber = eKey.RollingStartIntervalNumber
			rKey.RollingPeriod = eKey.RollingPeriod
			rKey.DaysSinceOnsetOfSymptoms = eKey.DaysSinceOnsetOfSymptoms
			restKeys = append(restKeys, *rKey)
		}
		if bindata {
			byteData := make([]byte, 0)
			for _, k := range restKeys {
				dk, err := base64.StdEncoding.DecodeString(k.ExpKey)
				if err == nil {
					byteData = append(byteData, dk...)
				}
				b := make([]byte, 4)
				b = unsafeCaseInt32ToBytes(k.RollingStartIntervalNumber)
				byteData = append(byteData, b...)
				b = unsafeCaseInt32ToBytes(k.RollingPeriod)
				byteData = append(byteData, b...)
				b = unsafeCaseInt32ToBytes(k.DaysSinceOnsetOfSymptoms)
				byteData = append(byteData, b...)
			}
			err := c.Send(byteData)
			if err != nil {
				log.Errorf("GetAllExpKeysByDate: %s", err.Error())
				return err
			}
		} else {
			err := c.SendString(createCSV(restKeys, false))
			if err != nil {
				log.Errorf("GetAllExpKeysByDate: %s", err.Error())
				return err
			}
		}

	}
	return nil
}

func PostExpKey(c *fiber.Ctx) error {
	log.Info("PostExpKey")
	auth := strings.TrimSpace(c.Get("Authorization"))
	if len(auth) <= 0 {
		log.Error("missing auth token")
		err := c.SendStatus(http.StatusForbidden)
		if err != nil {
			log.Errorf("PostExpKey: %s", err.Error())
			return err
		}
		return nil
	} else {
		log.Debugf("<PostExpKey> auth token: %s", auth)
		_, err := authtoken.IsValidToken(auth)
		if err != nil {
			err := fmt.Errorf("invalid token: %s", auth)
			log.Errorf(err.Error())
			return err
		}
	}

	data := c.Request().Body()
	if len(data)%28 != 0 {
		return fmt.Errorf("invalid number of bytes: %d", len(data))
	}

	today := time.Now().UTC()
	cDayStr := today.Format("2006-01-02")
	log.Debugf("<PostExpKey> currentDay: %s ", cDayStr)

	rawKeys := [][]byte{}
	for i := 0; i < len(data); i += 28 {
		rawKey := data[i : i+28]
		rawKeys = append(rawKeys, rawKey)
		expK := expkey.ExpKey{}
		key := base64.StdEncoding.EncodeToString(rawKey[0:16])
		expK.Day = cDayStr
		expK.ExpKey = key
		expK.RollingStartIntervalNumber = int32(rawKey[19])<<24 | int32(rawKey[18])<<16 | int32(rawKey[17])<<8 | int32(rawKey[16])
		expK.RollingPeriod = int32(rawKey[23])<<24 | int32(rawKey[22])<<16 | int32(rawKey[21])<<8 | int32(rawKey[20])
		expK.DaysSinceOnsetOfSymptoms = int32(rawKey[27])<<24 | int32(rawKey[26])<<16 | int32(rawKey[25])<<8 | int32(rawKey[24])
		expK.KeyProvider = expkey.KEYPROVIDER_PANTRA
		err := expkey.StoreExpKey(&expK)
		if err != nil {
			return fmt.Errorf("storing ExpKey failed for day: %s", cDayStr)
		}
	}

	log.Debugf("<PostExpKey> received keys: %d", len(rawKeys))

	_, err := authtoken.MarkUsed(auth)
	if err != nil {
		log.Error(err.Error())
		c.SendStatus(http.StatusInternalServerError)
		return err
	} else {
		err = c.SendString(fmt.Sprintf("OK, added keys %d", len(rawKeys)))
		if err != nil {
			log.Errorf("PostExpKey: %s", err.Error())
			return err
		} else {
			return nil
		}
	}
}

func GenerateTokens(c *fiber.Ctx) error {
	log.Info("GenerateTokens")
	nStr := c.Params("n")
	n, err := strconv.Atoi(nStr)
	if err == nil {
		authtoken.Generate(n)
		c.SendStatus(http.StatusOK)
		return nil
	} else {
		err := fmt.Errorf("invalid paramter: %s", nStr)
		log.Error(err.Error())
		return err
	}
}

func PopToken(c *fiber.Ctx) error {
	log.Info("popToken")

	auth := strings.TrimSpace(c.Get("Authorization"))
	if len(auth) <= 0 {
		log.Error("missing auth token")
		err := c.SendStatus(http.StatusForbidden)
		if err != nil {
			log.Errorf("<PopToken> %s", err.Error())
			return err
		}
		return nil
	} else {
		log.Debugf("<PopToken> auth token: %s", auth)
		if config.AccessToken != auth {
			err := fmt.Errorf("invalid token: %s", auth)
			log.Errorf(err.Error())
			return err
		}
	}

	token, err := authtoken.PopToken()
	if err != nil {
		return err
	} else {
		c.SendString(token)
		return nil
	}
}

func createCSV(restKeys []expkey.ExpKeyRest, withDate bool) string {
	resultCSV := ""
	for _, eKey := range restKeys {
		day := eKey.Day
		key := eKey.ExpKey
		rollStart := eKey.RollingStartIntervalNumber
		rollPer := eKey.RollingPeriod
		days := eKey.DaysSinceOnsetOfSymptoms
		line := fmt.Sprintf("%s,%d,%d,%d", key, rollStart, rollPer, days)
		if withDate {
			resultCSV = fmt.Sprintf("%s%s,%s\n", resultCSV, day, line)
		} else {
			resultCSV = fmt.Sprintf("%s%s\n", resultCSV, line)
		}
	}
	return resultCSV
}
