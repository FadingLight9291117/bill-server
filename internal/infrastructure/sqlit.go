package infrastructure

import (
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

	return db, nil
}
