package handlers

import (
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
)

// int is a placeholder for the database type
func GetJournalAccessPerTime(db *gorm.DB, w http.ResponseWriter, req *http.Request) {
	// Example code once the database connection is made:
	//
	// employees := []model.Employee{}
	// db.Find(&employees)
	// respondJSON(w, http.StatusOK, employees)

	io.WriteString(w, "Hello from JournalAccessPerTime route")
}
