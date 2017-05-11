package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
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
	Owner     *User `gorm:"ForeignKey:owner_id"`
}

// Project model stores the database structure
type Project struct {
	gorm.Model
	Name        string
	Description string
	Tasks       []Task
}

// Task model stores the task related structure
type Task struct {
	gorm.Model
	Title       string
	Description string
	ParentID    int    `gorm:"index"`
	Tasks       []Task `gorm:"ForeignKey:parent_id"`
	Parent      *Task  `gorm:"ForeignKey:parent_id"`
}

func main() {
	fmt.Println("vim-go")
	db, err := gorm.Open("postgres", "postgres://localhost/dealgrok_dev?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Project{}, &Task{})
	r := mux.NewRouter()
	r.HandleFunc("/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		encoder := json.NewEncoder(rw)
		rw.Header().Add("Content-Type", "application/json")
		encoder.Encode(map[string]string{"greeting": fmt.Sprintf("Hello %s", vars["name"])})
	})
	http.ListenAndServe(":3000", r)
}
