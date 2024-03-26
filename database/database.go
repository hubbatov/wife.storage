package database

import (
	"fmt"
	"log"

	"wife.storage/errors"
	"wife.storage/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //pg adapter
)

// DBManager is a standalone DatabaseManager object
var DBManager DatabaseManager

// DatabaseManager (holds database connection)
type DatabaseManager struct {
	DataBase *gorm.DB
}

// CreateDb creates new DatabaseManager
func CreateDb(config *models.DatabaseConfig) error {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	log.Println("Connecting to database:", dbinfo)

	db, err := gorm.Open(config.Provider, dbinfo)

	errors.HandleError(errors.ConvertCustomError(err))

	if err == nil {
		log.Println("Migrate ...")
		db.AutoMigrate(&models.User{})

		DBManager.DataBase = db

		log.Println("Database is ready.")
		return nil
	}

	return err
}
