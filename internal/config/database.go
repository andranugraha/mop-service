package config

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var globalDB *gorm.DB

func initDatabase() error {
	connectionURL := dbURL
	db, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return errors.New("failed to connect to database: " + err.Error())
	}

	dbSQL, err := db.DB()
	if err != nil {
		return errors.New("failed to connect to database: " + err.Error())
	}

	dbSQL.SetMaxIdleConns(20)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)
	dbSQL.SetConnMaxIdleTime(15 * time.Minute)

	globalDB = db
	return nil
}

func GetDB() *gorm.DB {
	return globalDB
}
