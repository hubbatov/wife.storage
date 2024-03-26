package database

import (
	"wife.storage/models"
	"wife.storage/rest"
)

func (d *DatabaseManager) Users() []models.User {
	var table []models.User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) User(userLogin, userPassword string) []models.User {
	var table []models.User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) CreateUser(userdata rest.User) []error {
	u := models.CreateUser(userdata)
	return d.DataBase.Create(&u).GetErrors()
}
