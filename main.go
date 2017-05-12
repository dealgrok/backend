package main

import (
	"encoding/json"
	"fmt"
	"github.com/dealgrok/backend/config"
	"github.com/dealgrok/backend/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	config := config.Init()
	db, err := models.Init(config)
	if err != nil {
		log.Fatal(err)
	}
	models.Migrate(db)
	r := mux.NewRouter()
	r.HandleFunc("/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		encoder := json.NewEncoder(rw)
		rw.Header().Add("Content-Type", "application/json")
		err := encoder.Encode(map[string]string{"greeting": fmt.Sprintf("Hello %s", vars["name"])})
		if err != nil {
			panic(err)
		}
	})
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
