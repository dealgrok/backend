package testutils

import (
	"github.com/dealgrok/backend/config"
	"github.com/dealgrok/backend/models"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupTestDB() (*gorm.DB, error) {
	os.Setenv("DATABASE_URL", "postgresql://localhost/dealgrok_test?sslmode=disable")
	c := config.Init()
	db, err := models.Init(c)
	if NoErr("unable to initialize database", err) {
		models.Migrate(db)
		return db.Begin(), nil
	}
	return nil, err
}

func NoErr(reason string, err error) bool {
	if err != nil {
		log.Fatalf("%s - %s", reason, err)
		return false
	} else {
		return true
	}
}
