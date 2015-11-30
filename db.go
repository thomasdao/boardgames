package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// SetupDB connection
func SetupDB() {
	newDB, err := gorm.Open("postgres", "user=thomasdao dbname=thomasdao sslmode=disable")
	if err != nil {
		panic(err)
	}

	db = &newDB
}
