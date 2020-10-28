package updaterservice

import (
	"encoding/base64"
	configUtil "github.com/pantraeu/pantra_server/pkg/pantra_server/confutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/expkeyutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"time"
)

func UpdateExpKeys() {
	config := configUtil.GetConfig()
	dataRoot := config.DataPath

	log.Info("check for missing keys")
	dataFiles := expkeyutil.GetBinFiles(dataRoot)
	for _, df := range dataFiles {
		keys, err := expkeyutil.GetTmpKey(df)
		if err == nil {

			fromT := time.Unix(int64(*keys.StartTimestamp), 0)
			//toT := time.Unix(int64(*keys.EndTimestamp), 0)
			cDayStr := fromT.Format("2006-01-02")
			cDayShortStr := fromT.Format("20060102")
			curKeys, err := expkey.GetExpKeysByDate(cDayShortStr, 0, 1)
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
					ek.KeyProvider = expkey.KEYPROVIDER_RKI

					expKeys = append(expKeys, *ek)

					if len(expKeys) >= config.InsertBatchSize {
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
		} else {
			log.Errorf("fetching ExpKeys failed: %s", err.Error())
		}
	}
}
