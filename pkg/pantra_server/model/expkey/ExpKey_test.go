package expkey

import (
	database "github.com/dermicha/goutils/database"
	utils2 "github.com/gofiber/fiber/v2/utils"
	log "github.com/sirupsen/logrus"
	"strings"
	"testing"
)

var (
	testDbName = ":memory:"
	//dbName     = "testdatabase"
	//testDbName = fmt.Sprintf("%s_test", dbName)
)

func setup() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)

	log.Info("test setup")
	database.CleanUpDb(testDbName)
	database.InitDatabase(testDbName)
	database.MigrateDatabase(&ExpKey{})
}

func tearDown() {
	log.Info("test teardown")
	database.CloseDatabase()
	database.CleanUpDb(testDbName)
}

func TestCrud(t *testing.T) {
	setup()
	defer tearDown()

	tObj := ExpKey{}
	tObj.Day = "2020-10-03"
	tObj.ExpKey = utils2.UUID()

	resCreate := database.GetDb().Create(&tObj)
	if resCreate.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resCreate.RowsAffected)
	}
	if tObj.ID != 1 {
		t.Fatalf("wrong object id: %d", tObj.ID)
	}

	tObj2 := ExpKey{}
	tObj2.Day = "2020-10-03"
	tObj2.ExpKey = utils2.UUID()

	resCreate2 := database.GetDb().Create(&tObj2)
	if resCreate2.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resCreate2.RowsAffected)
	}
	if tObj2.ID != 2 {
		t.Fatalf("wrong object id: %d", tObj2.ID)
	}

	resDel := database.GetDb().Delete(tObj)
	if resDel.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resDel.RowsAffected)
	}

	resDel2 := database.GetDb().Delete(tObj2)
	if resDel2.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resDel2.RowsAffected)
	}
}

func TestUniqueIndex(t *testing.T) {
	setup()
	defer tearDown()

	tObj := ExpKey{}
	tObj.Day = "2020-10-03"
	tObj.ExpKey = utils2.UUID()
	resCreate1 := database.GetDb().Create(&tObj)
	resCreate2 := database.GetDb().Create(&tObj)

	if resCreate1.RowsAffected != 1 {
		t.Fatalf("wrong number of rows effected: %d", resCreate1.RowsAffected)
	}
	if tObj.ID != 1 {
		t.Fatalf("wrong object id: %d", tObj.ID)
	}
	if resCreate2.RowsAffected != 0 {
		t.Fatalf("wrong number of rows effected: %d", resCreate2.RowsAffected)
	}
	if !strings.Contains(resCreate2.Error.Error(), "UNIQUE constraint failed:") {
		t.Fatalf("unique index failed")
	}
}
