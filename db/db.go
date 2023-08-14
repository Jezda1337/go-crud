package db

import (
	"github.com/Jezda1337/go-crud/domain/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := "host=db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Belgrade"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = autoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&book.Book{},
	)

	if err != nil {
		return err
	}

	return nil
}
