package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/rest_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Panic("Error while connecting to database")
	}

	// db.DB().SetMaxIdleConns(4)

	return db
}
