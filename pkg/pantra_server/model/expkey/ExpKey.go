package expkey

import (
	"fmt"
	database "github.com/dermicha/goutils/database_pg"
	"gorm.io/gorm"
	"time"
)

const (
	KEYPROVIDER_RKI    = "kp_rki"
	KEYPROVIDER_PANTRA = "kp_pantra"
)

type ExpKey struct {
	gorm.Model
	Day                        string `gorm:"column:day"`
	ExpKey                     string `gorm:"index:idx_expkeys,unique,column:exp_key"`
	RollingStartIntervalNumber int32  `gorm:"column:interval_number"`
	RollingPeriod              int32  `gorm:"column:rolling_period"`
	DaysSinceOnsetOfSymptoms   int32  `gorm:"column:days_since"`
	KeyProvider                string `gorm:"column:key_provider"`
	//TransmissionRiskLevel      int32  `gorm:"column:transmission_risk_level"`
}

type ExpKeyRest struct {
	Day                        string `json:"day"`
	ExpKey                     string `json:"exp_key"`
	RollingStartIntervalNumber int32  `json:"rolling_start_interval_number"`
	RollingPeriod              int32  `json:"rolling_period"`
	DaysSinceOnsetOfSymptoms   int32  `json:"days_since_onset_of_symptoms"`
	//TransmissionRiskLevel      int32  `gorm:"column:transmission_risk_level"`
}

func GetExpKeysByOffset(offset int, page int, size int) ([]ExpKey, error) {
	if offset <= 13 && offset >= 0 {
		if page >= 0 {

			dbCon := database.GetDb()
			dbCon.Exec("set client_encoding to 'UTF8'")
			today := time.Now().UTC()
			dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", (offset * 24)))
			currentDay := today.Add(dur).Format("2006-01-02")

			var expKeys []ExpKey
			dbCon.
				Limit(size).
				Offset(page*size).
				Where("Day = ?", currentDay).
				Order("rolling_start_interval_number asc").
				Find(&expKeys)
			return expKeys, nil
		} else {
			return nil, fmt.Errorf("invalid page (>= 0): %d", page)
		}
	} else {
		return nil, fmt.Errorf("invalid offset (0-13): %d", offset)
	}
}

func GetExpKeysByDateByProvider(dateStr string, provider string, page int, size int) ([]ExpKey, error) {

	if page >= 0 {

		dbCon := database.GetDb()
		dbCon.Exec("set client_encoding to 'UTF8'")
		date, err := time.Parse("20060102", dateStr)
		if err != nil {
			return nil, fmt.Errorf("invalid date: %s", dateStr)
		}
		currentDay := date.Format("2006-01-02")
		var expKeys []ExpKey
		dbCon.
			Limit(size).
			Offset(page*size).
			Where("day = ?", currentDay).
			Where("key_provider = ?", provider).
			Order("exp_key asc").
			Find(&expKeys)
		return expKeys, nil
	} else {
		return nil, fmt.Errorf("invalid page (>= 0): %d", page)
	}
}

func GetAllExpKeysByDate(dateStr string, page int, size int) ([]ExpKey, error) {

	if page >= 0 {

		dbCon := database.GetDb()
		dbCon.Exec("set client_encoding to 'UTF8'")
		date, err := time.Parse("20060102", dateStr)
		if err != nil {
			return nil, fmt.Errorf("invalid date: %s", dateStr)
		}
		currentDay := date.Format("2006-01-02")
		var expKeys []ExpKey
		dbCon.
			Limit(size).
			Offset(page*size).
			Where("Day = ?", currentDay).
			Order("exp_key asc").
			Find(&expKeys)
		return expKeys, nil
	} else {
		return nil, fmt.Errorf("invalid page (>= 0): %d", page)
	}
}

func StoreExpKey(expkey *ExpKey) error {
	dbCon := database.GetDb()
	dbCon.Exec("set client_encoding to 'UTF8'")
	dbCon.Create(&expkey)
	return nil
}

func StoreExpKeys(expkey *[]ExpKey) error {
	dbCon := database.GetDb()
	dbCon.Exec("set client_encoding to 'UTF8'")
	dbCon.Create(&expkey)
	return nil
}
