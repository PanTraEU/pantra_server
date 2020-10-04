package main

import (
	"fmt"
	"github.com/dermicha/goutils/database"
	"github.com/gofiber/fiber/v2"
	utils2 "github.com/gofiber/fiber/v2/utils"
	configUtil "github.com/pantraeu/pantra_server/pkg/pantra_server/confutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/expkeyservice"
	expkey "github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"time"
)

func setupRoutes(app *fiber.App) {
	log.Println("setupRoutes")

	apiV1 := app.Group("/pantraserver/api/v1")

	apiV1.Get("/expkey/:offset/:page", expkeyservice.GetExpKeys)
	apiV1.Get("/expkey/:offset/:page/:size", expkeyservice.GetExpKeys)

}

func initData() {
	expKeys, _ := expkey.GetAllExpKeys()
	if len(expKeys) == 0 {
		for i := 0; i < 14; i++ { // start of the execution block
			today := time.Now().UTC()
			for j := 0; j < 100; j++ { // start of the execution block
				dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", (i * 24)))
				currentDay := today.Add(dur)
				log.Info(currentDay)
				ek := new(expkey.ExpKey)
				ek.ExpKey = utils2.UUID()
				ek.Day = currentDay.Format("2006-01-02")
				expkey.StoreExpKey(ek)
			}
		}
	}
}

func main() {
	log.Println("Welcome!")
	config := configUtil.GetConfig()
	database.InitDatabase(config.DbPath)
	database.MigrateDatabase(&expkey.ExpKey{})

	initData()

	app := fiber.New()
	setupRoutes(app)

	//defer database.DBConn.Close()
	err := app.Listen(":3000")
	if err != nil {
		log.Error("could not start server: ", err.Error())
	}
}
