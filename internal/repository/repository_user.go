package repository

import (
	"wife/internal/user"
	"wife/utils"
)

func (d *Repository) Users() []user.User {
	var table []user.User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *Repository) User(userLogin, userPassword string) []user.User {
	var table []user.User
	d.DataBase.Order("id").Find(&table)
	return table
}

func (d *Repository) Exists(login string, password string) bool {
	var user user.User
	d.DataBase.Where("login = ?", login).First(&user)

	err := utils.CheckPassword(password, user.Password)

	return err == nil
}

func (d *Repository) CreateUser(dto user.UserDto) []error {
	u, err := user.FromDto(dto)
	if err != nil {
		return []error{err}
	}
	return d.DataBase.Create(&u).GetErrors()
}
