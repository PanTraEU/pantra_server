package authtoken

import (
	"encoding/base64"
	"fmt"
	database "github.com/dermicha/goutils/database_pg"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AutToken struct {
	gorm.Model
	AuthToken    string `gorm:"index:idx_authtoken,unique,column:auth_token"`
	ValidState   bool   `gorm:"column:valid_state"`
	UserAssigned bool   `gorm:"column:user_assigned"`
	Permament    bool   `gorm:"column:permament"`
}

func IsValidToken(token string) (bool, error) {
	db := database.GetDb()

	ats := []AutToken{}

	db.Limit(1).
		Where("auth_token = ?", token).
		Where("valid_state = ?", true).
		Where("user_assigned = ?", true).
		Order("created_at asc").
		Find(&ats)

	if len(ats) == 1 {
		return true, nil
	} else {
		err := fmt.Errorf("not a vaild token: %s", token)
		return false, err
	}
}

func MarkUsed(token string) (bool, error) {
	db := database.GetDb()

	_, err := IsValidToken(token)
	if err != nil {
		return false, err
	} else {
		ats := []AutToken{}

		db.Limit(1).
			Where("auth_token = ?", token).
			Where("valid_state = ?", true).
			Where("user_assigned = ?", true).
			Order("created_at asc").
			Find(&ats)
		if len(ats) == 1 {
			at := ats[0]

			if !at.Permament {
				at.ValidState = false
				db.Save(at)
			}

			return true, nil
		} else {
			err := fmt.Errorf("not a vaild token: %s", token)
			return false, err
		}
	}
}

func PopToken() (string, error) {
	db := database.GetDb()

	ats := []AutToken{}

	db.Limit(1).
		Where("valid_state = ?", true).
		Where("user_assigned = ?", false).
		Order("created_at asc").
		Find(&ats)

	if len(ats) > 0 {
		at := ats[0]
		at.UserAssigned = true
		db.Save(at)

		return at.AuthToken, nil
	} else {
		err := fmt.Errorf("no more valid tokens")
		log.Error(err.Error())
		return "", err
	}
}

func Generate(n int) error {

	for i := 0; i < n; i++ {
		at := AutToken{}
		uid, _ := uuid.NewV4()
		rndStr := base64.StdEncoding.EncodeToString(uid.Bytes())[0:10]

		at.AuthToken = rndStr
		at.ValidState = true
		at.UserAssigned = false
		at.Permament = false

		database.GetDb().Create(&at)
	}
	return nil
}
