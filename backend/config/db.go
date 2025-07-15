package config

import (
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	SqliteDbPath = "./backend/db/voucher.db"
)

var DB *gorm.DB

func InitSQLiteDB() *gorm.DB {
	dbPath := "test.db"
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to SQLite database")
		log.Fatal("connection error:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	return db
}
