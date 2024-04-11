package user

import (
	"fmt"
	"log"

	"wife/configs"
	"wife/internal/errors"
	"wife/utils"

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
		db.AutoMigrate(&User{})

		log.Println("Database is ready.")
		return &Repository{
			DataBase: db,
		}, nil
	}

	return nil, err
}

func (d *Repository) Users() []User {
	var table []User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *Repository) User(userLogin, userPassword string) []User {
	var table []User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *Repository) Exists(login string, password string) bool {
	var user User
	d.DataBase.Where("login = ?", login).First(&user)

	err := utils.CheckPassword(password, user.Password)

	return err == nil
}

func (d *Repository) CreateUser(dto UserDto) []error {
	u, err := FromDto(dto)
	if err != nil {
		return []error{err}
	}
	return d.DataBase.Create(&u).GetErrors()
}
