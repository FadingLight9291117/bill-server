package infrastructure

import (
	"bill-go-fiber/internal/bill"
	"bill-go-fiber/internal/label"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToSqlite() (*gorm.DB, error) {
	dsn := "bill.db"
	//dsn := "file::memory:?cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.\n")
	}

	err = db.AutoMigrate(&label.Label{})
	if err != nil {
		panic("failed to migrate database.\n")
	}
	err = db.AutoMigrate(&bill.Bill{})
	if err != nil {
		panic("failed to migrate database.\n")
	}

	db.Create(&label.Label{ID: 1, Name: "bill"})

	return db, nil
}
