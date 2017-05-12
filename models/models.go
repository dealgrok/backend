package models

import (
	"github.com/dealgrok/backend/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"time"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique_index"`
	Password     string
	LastLoggedIn time.Time
}

type Organization struct {
	gorm.Model
	Name      string
	SubDomain string
	OwnerID   int
	Owner     User `gorm:"ForeignKey:owner_id"`
}

// Project model stores the database structure
type Project struct {
	gorm.Model
	Name        string
	Description string
	Tasks       []Task
	CreatedByID int
	CreatedBy   User
}

// Task model stores the task related structure
type Task struct {
	gorm.Model
	Title       string
	Description string
	ParentID    int    `gorm:"index"`
	Tasks       []Task `gorm:"ForeignKey:parent_id"`
	Parent      *Task
	CreatedByID int
	CreatedBy   User
}

// Migrate sets up the database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Organization{},
		&Project{},
		&Task{},
	)
}

func Init(c config.Config) (*gorm.DB, error) {
	log.Infof("Connecting to database %s", c.DatabaseUrl)
	return gorm.Open("postgres", c.DatabaseUrl)
}
