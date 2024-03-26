package database

import (
	"fmt"

	"wife.storage/errors"
	"wife.storage/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //pg adapter
)

const (
	//DbHost - database host
	DbHost = "localhost"
	//DbPort - database port
	DbPort = 5432
	//DbUser - database user
	DbUser = "postgres"
	//DbPassword - database password
	DbPassword = "123456"
	//DbName - database db name
	DbName = "smarthouse"
)

// DBManager is a standalone DatabaseManager object
var DBManager DatabaseManager

// DatabaseManager (holds database connection)
type DatabaseManager struct {
	DataBase *gorm.DB
}

// CreateDb creates new DatabaseManager
func CreateDb() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err := gorm.Open("postgres", dbinfo)

	errors.HandleError(errors.ConvertCustomError(err))

	db.AutoMigrate(&models.User{})

	DBManager.DataBase = db
}
