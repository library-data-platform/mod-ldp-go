package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type LogObj struct {
	LogTime     string  `json:"log_time"`
	Tablename   string  `gorm:"column:table_name" json:"table_name"`
	ElapsedTime float64 `json:"elapsed_time"`
}

func (c LogObj) TableName() string {
	return "ldpsystem.log"
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&LogObj{})
	return db
}
