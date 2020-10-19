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
			rKey.RollingStartIntervalNumber = eKey.RollingStartIntervalNumber
			rKey.RollingPeriod = eKey.RollingPeriod
			rKey.DaysSinceOnsetOfSymptoms = eKey.DaysSinceOnsetOfSymptoms
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
