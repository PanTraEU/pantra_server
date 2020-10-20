package updaterservice

import (
	"encoding/base64"
	"fmt"
	utils2 "github.com/gofiber/fiber/v2/utils"
	configUtil "github.com/pantraeu/pantra_server/pkg/pantra_server/confutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/expkeyutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"time"
)

func _UpdateExpKeys() {
	log.Info("check for missing keys")
	for i := 0; i < 14; i++ { // start of the execution block
		today := time.Now().UTC()
		dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", (i * 24)))
		cDay := today.Add(dur)
		cDayShortStr := cDay.Format("20060102")
		cDayStr := cDay.Format("2006-01-02")
		curKeys, err := expkey.GetExpKeysByDate(cDayShortStr, 0, 10)
		if err != nil || len(curKeys) == 0 {
			log.Info("add keys for day: ", cDayStr)
			for j := 0; j < 100; j++ { // start of the execution block
				ek := new(expkey.ExpKey)
				ek.ExpKey = utils2.UUID()
				ek.Day = cDayStr
				err := expkey.StoreExpKey(ek)
				if err != nil {
					log.Panic("ExpKey could not be stored: ", err.Error())
				}
			}
		}
	}
}

func UpdateExpKeys() {
	config := configUtil.GetConfig()
	dataRoot := config.DataPath

	log.Info("check for missing keys")
	dataFiles := expkeyutil.GetBinFiles(dataRoot)
	for _, df := range dataFiles {
		keys, err := expkeyutil.GetTmpKey(df)
		fromT := time.Unix(int64(*keys.StartTimestamp), 0)
		//toT := time.Unix(int64(*keys.EndTimestamp), 0)
		cDayStr := fromT.Format("2006-01-02")
		cDayShortStr := fromT.Format("20060102")
		curKeys, err := expkey.GetExpKeysByDate(cDayShortStr, 0, 10)
		if err != nil || len(curKeys) == 0 {
			log.Infof("add %d keys for day: %s", len(keys.Keys), cDayStr)
			expKeys := make([]expkey.ExpKey, 0)
			for _, k := range keys.Keys {
				hexKey := base64.StdEncoding.EncodeToString(k.KeyData)
				rollStart := k.RollingStartIntervalNumber
				rollPer := k.RollingPeriod
				days := k.DaysSinceOnsetOfSymptoms
				ek := new(expkey.ExpKey)
				ek.Day = cDayStr
				ek.ExpKey = hexKey
				if rollPer != nil {
					ek.RollingPeriod = *rollPer
				} else {
					ek.RollingPeriod = 0
				}
				if rollStart != nil {
					ek.RollingStartIntervalNumber = *rollStart
				} else {
					ek.RollingStartIntervalNumber = 0
				}
				if days != nil {
					ek.DaysSinceOnsetOfSymptoms = *days
				} else {
					ek.DaysSinceOnsetOfSymptoms = 0
				}
				expKeys = append(expKeys, *ek)

				if len(expKeys) >= 500 {
					log.Infof("adding batch of size %d", len(expKeys))
					err := expkey.StoreExpKeys(&expKeys)
					if err != nil {
						log.Panic("ExpKey could not be stored: ", err.Error())
					}
					expKeys = make([]expkey.ExpKey, 0)
				}
			}
			err := expkey.StoreExpKeys(&expKeys)
			if err != nil {
				log.Panic("ExpKey could not be stored: ", err.Error())
			}
		}
	}
}
