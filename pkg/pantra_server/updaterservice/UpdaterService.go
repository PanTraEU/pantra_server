package updaterservice

import (
	"fmt"
	utils2 "github.com/gofiber/fiber/v2/utils"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"time"
)

func UpdateExpKeys() {
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
				expkey.StoreExpKey(ek)
			}
		}
	}
}
