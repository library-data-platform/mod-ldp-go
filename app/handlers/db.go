package handlers

import (
	"fmt"
	"net/http"

	"github.com/folio-org/mod-ldp/app/model"
	"github.com/jinzhu/gorm"
)

func GetDbStatus(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	if err := db.DB().Ping(); err != nil {
		fmt.Println(err)
		respondJSON(w, http.StatusOK, map[string]string{"status": "error"})
	} else {
		respondJSON(w, http.StatusOK, map[string]string{"status": "connected"})
	}
}

func GetLogTable(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	logs := []model.LogObj{}
	db.Find(&logs)
	respondJSON(w, http.StatusOK, logs)
}
