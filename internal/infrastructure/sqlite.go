package infrastructure

import (
	"bill-go-fiber/internal/bill"
	"bill-go-fiber/internal/label"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToSqlite() (*gorm.DB, error) {
	var dsn string
	switch os.Getenv("ENV") {
	case "prod":
		dsn = "bill.db"
	default:
		dsn = "file::memory:?cache=shared"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.\n")
	} else {
		log.Println("Connecting to sqlite.")
	}

	err = db.AutoMigrate(&label.Label{})
	if err != nil {
		panic("failed to migrate database.\n")
	}
	err = db.AutoMigrate(&bill.Bill{})
	if err != nil {
		panic("failed to migrate database.\n")
	}

	// migrate data
	if os.Getenv("ENV") != "prod" {
		go func() {
			sql, err := os.ReadFile("scripts/sqliteMigrations.sql")
			if err != nil {
				//panic("failed to read migration sql file: " + err.Error())
				log.Fatal("failed to read migration sql file: " + err.Error())
			}
			res := db.Exec(string(sql))
			if res.Error != nil {
				log.Fatal("failed to migrate" + res.Error.Error())
			}
			log.Println("SQL migrated successfully!")
		}()
	}

	return db, nil
}
