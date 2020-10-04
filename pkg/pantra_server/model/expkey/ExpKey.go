package expkey

import (
	"fmt"
	"github.com/dermicha/goutils/database"
	"gorm.io/gorm"
	"time"
)

type ExpKey struct {
	gorm.Model
	Day    string `gorm:"column:day"`
	ExpKey string `gorm:"column:exp_key" gorm:"uniqueIndex"`
}

type ExpKeyRest struct {
	Day    string `json:"day"`
	ExpKey string `json:"exp_key"`
}

func GetAllExpKeys() ([]ExpKey, error) {
	dbCon := database.GetDb()
	var expKeys []ExpKey
	dbCon.Find(&expKeys)
	return expKeys, nil
}

func GetExpKeys(offset int, page int, size int) ([]ExpKey, error) {
	if offset <= 13 && offset >= 0 {
		if page >= 0 {

			dbCon := database.GetDb()

			today := time.Now().UTC()
			dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", (offset * 24)))
			currentDay := today.Add(dur).Format("2006-01-02")

			var expKeys []ExpKey
			dbCon.
				Limit(size).
				Offset(page*size).
				Where("Day = ?", currentDay).
				Find(&expKeys)
			return expKeys, nil
		} else {
			return nil, fmt.Errorf("invalid page (>= 0): %d", page)
		}
	} else {
		return nil, fmt.Errorf("invalid offset (0-13): %d", offset)
	}
}

func StoreExpKey(expkey *ExpKey) error {
	dbCon := database.GetDb()
	dbCon.Create(&expkey)
	return nil
}
