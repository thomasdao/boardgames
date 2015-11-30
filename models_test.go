package main_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/thomasdao/boardgames"
)

var db *gorm.DB

func setup() {
	newDB, err := gorm.Open("postgres", "user=thomasdao dbname=thomasdao sslmode=disable")
	if err != nil {
		panic(err)
	}

	db = &newDB
	// Enable Logger
	db.LogMode(true)

}

func tearDown() {
	if db != nil {
		db.Close()
	}
}

func TestLoad(t *testing.T) {
	setup()
	defer tearDown()

	err := db.DB().Ping()
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(
		&main.User{},
		&main.Game{},
		&main.Rating{},
		&main.Similarity{},
	)

	var user main.User
	err = db.First(&user).Error
	if err != nil {
		t.Fatal(err)
	}

	var rating main.Rating
	err = db.Preload("User").Preload("Game").First(&rating).Error
	if err != nil {
		t.Fatal(err)
	}

	var sim main.Similarity
	err = db.Preload("GameA").Preload("GameB").First(&sim).Error
	if err != nil {
		t.Fatal(err)
	}

	var game main.Game
	err = db.First(&game).Error
	if err != nil {
		t.Fatal(err)
	}

}
