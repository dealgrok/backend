package models_test

import (
	"fmt"
	"github.com/dealgrok/backend/models"
	"github.com/dealgrok/backend/testutils"
	"github.com/jinzhu/gorm"
	"testing"
)

var db *gorm.DB

func TestGetChildTasks(t *testing.T) {
	fmt.Println(db)
	var tasks []models.Task
	db.Preload("Parent").Where("parent_id=1").Find(&tasks)
	fmt.Print(tasks[0].Parent)
	if len(tasks) != 2 {
		t.Error("Exepected 3 tasks")
	}
}

func TestMain(m *testing.M) {
	var err error
	db, err = testutils.SetupTestDB()
	if testutils.NoErr("unable to setup db", err) {
		m.Run()
		db.Rollback()
	}
}
