package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize database connection
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("ERROR INITIALIZING DATABASE CONFIG: %w", err)

	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize Logger
	logger = NewLogger(p)
	return logger
}
