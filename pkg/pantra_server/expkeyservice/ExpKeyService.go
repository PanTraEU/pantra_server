package expkeyservice

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	expkey "github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	"strconv"
	"time"
)

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
		c.SendStatus(404)
	} else {
		restKeys := []expkey.ExpKeyRest{}
		for _, eKey := range expKeys {
			rKey := new(expkey.ExpKeyRest)
			rKey.Day = eKey.Day
			rKey.ExpKey = eKey.ExpKey
			restKeys = append(restKeys, *rKey)
		}
		c.SendString(createCSV(restKeys, true))
	}
	return nil
}

func GetExpKeysByDate(c *fiber.Ctx) error {
	today := time.Now().UTC()
	currentDay := today.Format("2006-01-02")

	dateStr := c.Params("date", currentDay)
	if len(dateStr) != 8 {
		return fmt.Errorf("invalid date (yyyymmdd): ", dateStr)
	}

	page, err := strconv.Atoi(c.Params("page", "0"))
	if err != nil {
		return err
	}

	size, err := strconv.Atoi(c.Params("size", "10"))
	if err != nil {
		return err
	}

	expKeys, err := expkey.GetExpKeysByDate(dateStr, page, size)
	if err != nil {
		return err
	}

	if len(expKeys) == 0 {
		c.SendStatus(404)
	} else {
		restKeys := []expkey.ExpKeyRest{}
		for _, eKey := range expKeys {
			rKey := new(expkey.ExpKeyRest)
			rKey.Day = eKey.Day
			rKey.ExpKey = eKey.ExpKey
			restKeys = append(restKeys, *rKey)
		}
		c.SendString(createCSV(restKeys, false))
	}
	return nil
}

func createCSV(restKeys []expkey.ExpKeyRest, withDate bool) string {
	resultCSV := ""
	for _, eKey := range restKeys {
		day := eKey.Day
		key := eKey.ExpKey
		if withDate {
			resultCSV = fmt.Sprintf("%s%s,%s\n", resultCSV, day, key)
		} else {
			resultCSV = fmt.Sprintf("%s%s\n", resultCSV, key)
		}
	}
	return resultCSV
}

//func GetExpKey(c *fiber.Ctx) {
//	id := c.Params("id")
//	dbCon := database.GetDb()
//	var expKey expkey.ExpKey
//	dbCon.Find(&expKey, id)
//	c.JSON(expKey)
//}
//
//func NewExpKey(c *fiber.Ctx) error {
//	dbCon := database.GetDb()
//	expKey := new(expkey.ExpKey)
//	if err := c.BodyParser(expKey); err != nil {
//		c.Status(503).SendString(err.Error())
//		return fmt.Errorf("invalid json")
//	}
//
//	dbCon.Create(&expKey)
//	c.JSON(expKey)
//	return nil
//}
//
//func DeleteExpKey(c *fiber.Ctx) error {
//	id := c.Params("id")
//	dbCon := database.GetDb()
//
//	var expKey expkey.ExpKey
//	dbCon.First(&expKey, id)
//	if expKey.ExpKey == "" {
//		c.Status(500).SendString("No ExpKey Found with ID")
//		return fmt.Errorf("ExpKey not found with id: %d", id)
//	}
//	dbCon.Delete(&expKey)
//	c.SendString("ExpKey Successfully deleted")
//	return nil
//}
