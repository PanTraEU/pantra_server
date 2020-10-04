package expkeyservice_test

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func setup() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)

	log.Info("test setup")
}

func tearDown() {
	log.Info("test teardown")
}

func sendReq(t *testing.T, app *fiber.App, req *http.Request, respRegexp []string, testDesc string) {
	res, err := app.Test(req, 100)
	assert.Nilf(t, err, testDesc)

	assert.Equalf(t, 200, res.StatusCode, testDesc)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nilf(t, err, testDesc)
	for _, regExp := range respRegexp {
		re := regexp.MustCompile(regExp)
		match := re.FindStringIndex(string(body))
		assert.Greater(t, len(match), 0)
	}

}

func SimpleTest(t *testing.T) {
	setup()
	defer tearDown()

	testDesc := "about test"
	app := fiber.New()
	req, _ := http.NewRequest(
		"POST",
		"/pantraserver/api/v1/expkey/0/0",
		nil,
	)

	sendReq(t, app, req, []string{}, testDesc)
}
