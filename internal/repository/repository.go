package repository

import (
	"fmt"
	"log"

	"wife/configs"
	"wife/internal/errors"
	"wife/internal/reminder"
	"wife/internal/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //pg adapter
)

// DatabaseManager (holds database connection)
type Repository struct {
	DataBase *gorm.DB
}

func NewRepository(config *configs.DatabaseConfig) (*Repository, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	log.Println("Connecting to database:", dbinfo)

	db, err := gorm.Open(config.Provider, dbinfo)

	errors.HandleError(errors.ConvertCustomError(err))

	if err == nil {
		log.Println("Migrate ...")
		db.AutoMigrate(&user.User{})
		db.AutoMigrate(&reminder.Reminder{})

		log.Println("Database is ready.")
		return &Repository{
			DataBase: db,
		}, nil
	}

	return nil, err
}
