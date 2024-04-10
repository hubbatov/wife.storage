package database

import (
	"fmt"
	"log"

	"wife/internal/api"
	"wife/internal/errors"

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
func CreateDb(config *DatabaseConfig) error {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	log.Println("Connecting to database:", dbinfo)

	db, err := gorm.Open(config.Provider, dbinfo)

	errors.HandleError(errors.ConvertCustomError(err))

	if err == nil {
		log.Println("Migrate ...")
		db.AutoMigrate(&User{})

		DBManager.DataBase = db

		log.Println("Database is ready.")
		return nil
	}

	return err
}

func (d *DatabaseManager) Users() []User {
	var table []User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) User(userLogin, userPassword string) []User {
	var table []User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) CreateUser(userdata api.User) []error {
	u := CreateUser(userdata)
	return d.DataBase.Create(&u).GetErrors()
}
