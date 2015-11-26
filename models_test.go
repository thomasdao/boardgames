package boardgames_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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
		&boardgames.User{},
		&boardgames.Game{},
		&boardgames.Rating{},
		&boardgames.Similarity{},
	)

	var user boardgames.User
	err = db.First(&user).Error
	if err != nil {
		t.Fatal(err)
	}

	var rating boardgames.Rating
	err = db.Preload("User").Preload("Game").First(&rating).Error
	if err != nil {
		t.Fatal(err)
	}

	var sim boardgames.Similarity
	err = db.Preload("GameA").Preload("GameB").First(&sim).Error
	if err != nil {
		t.Fatal(err)
	}

	var game boardgames.Game
	err = db.First(&game).Error
	if err != nil {
		t.Fatal(err)
	}

}
