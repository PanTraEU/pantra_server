package main

import (
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/dermicha/goutils/database"
	"github.com/gofiber/fiber/v2"
	utils2 "github.com/gofiber/fiber/v2/utils"
	_ "github.com/pantraeu/pantra_server/cmd/pantra_server/docs"
	configUtil "github.com/pantraeu/pantra_server/pkg/pantra_server/confutil"
	"github.com/pantraeu/pantra_server/pkg/pantra_server/expkeyservice"
	expkey "github.com/pantraeu/pantra_server/pkg/pantra_server/model/expkey"
	log "github.com/sirupsen/logrus"
	"time"
)

func setupRoutes(app *fiber.App) {
	log.Println("setupRoutes")

	app.Use("/swagger", swagger.Handler)

	apiV1 := app.Group("/pantraserver/api/v1")

	apiV1.Get("/expkey/bydate/:date/:page", GetExpKeysByDate)
	apiV1.Get("/expkey/bydate/:date/:page/:size", GetExpKeysByDate)

	apiV1.Get("/expkey/:offset/:page", GetExpKeysByOffset)
	apiV1.Get("/expkey/:offset/:page/:size", GetExpKeysByOffset)

}

// GetExpKeysByOffset godoc
// @Summary access ExposureKeys by date
// @Description access to ExposureKeys for recent 14 days
// @ID get-expkeys-by-offset
// @Tags ExpKey
// @Accept  text/plain
// @Produce text/plain
// @Param offset path int true "an index in days beginning with today, vaild from today up to 13 backwards (0-13)"
// @Param page path int true "selects the batch of ExposureKeys for selected day, 0-n, return HTTP Status 404 if no more keys are available"
// @Param size path int false "defines amount of ExposureKeys per request, default is 10"
// @Success 200 {string} string
// @Failure 404 {String} string
// @Router /v1/expkey/{offset}/{page}/{size} [get]
func GetExpKeysByOffset(c *fiber.Ctx) error {
	return expkeyservice.GetExpKeysByOffset(c)
}

// GetExpKeysByDate godoc
// @Summary access ExposureKeys by date
// @Description access to ExposureKeys for recent 14 days
// @ID get-expkeys-by-date
// @Tags ExpKey
// @Accept  text/plain
// @Produce text/plain
// @Param date path string true "day as yyyymmdd for which ExposureKeys are requested, vaild from today up to 13 backwards (UTC based)"
// @Param page path int true "selects the batch of ExposureKeys for selected day, 0-n, return HTTP Status 404 if no more keys are available"
// @Param size path int false "defines amount of ExposureKeys per request, default is 10"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Router /v1/expkey/bydate/{date}/{page}/{size} [get]
func GetExpKeysByDate(c *fiber.Ctx) error {
	return expkeyservice.GetExpKeysByDate(c)
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

// @title PanTra Server API
// @version 0.1
// @description Swagger API docs
// @termsOfService http://swagger.io/terms/
// @contact.name derMicha
// @contact.email dermicha@dermicha.de
// @license.name GNU GENERAL PUBLIC LICENSE Version 3.0
// @license.url https://www.gnu.org/licenses/
// @host http://mqtt.pantra.eu
// @BasePath /pantraserver/api
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
