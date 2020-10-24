package expkeyservice

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	expkey "github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"unsafe"
)

const BYTES_IN_INT32 = 4

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
		return err
	}

	size, err := strconv.Atoi(c.Params("size", "10"))
	if err != nil {
		return err
	}

	expKeys, err := expkey.GetExpKeysByOffset(offset, page, size)
	if err != nil {
		return err
	}

	if len(expKeys) == 0 {
		err = c.SendStatus(404)
		if err != nil {
			log.Error(err)
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
			log.Error(err)
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
		return err
	}

	size, err := strconv.Atoi(c.Params("size", "10"))
	if err != nil {
		return err
	}

	log.Debugf("GetExpKeysByDate: %s / % d / %d", dateStr, page, size)

	expKeys, err := expkey.GetExpKeysByDate(dateStr, page, size)
	if err != nil {
		return err
	}

	if len(expKeys) == 0 {
		log.Error("GetExpKeysByDate: no more keys available")

		err := c.SendStatus(404)
		if err != nil {
			log.Error(err)
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
				log.Error(err)
			}
		} else {
			err := c.SendString(createCSV(restKeys, false))
			if err != nil {
				log.Error(err)
			}
		}

	}
	return nil
}

//func PostExpKeyByDate(c *fiber.Ctx, bindata bool) error {
//
//}

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
