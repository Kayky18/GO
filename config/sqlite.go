package config

import (
	"os"

	"github.com/kayky18/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	//Check if the database file exists
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating..")

		//Create the database path
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			logger.Error("ERRO CREATING DATABASE PATH")
			return nil, err
		}

		//Creating the database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("ERRO CREATING DATABASE FILE: %v", err)
			return nil, err
		}

		//Close the file
		file.Close()
	}

	// Open the SQLite database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("SQLITE Opening ERROR %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("AUTOMIGRATE ERROR: %v", err)
		return nil, err
	}

	return db, nil
}
