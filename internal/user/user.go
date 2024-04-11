package user

import (
	"time"
	"wife/utils"
)

type User struct {
	ID       int       `gorm:"primary_key" json:"id"`
	Name     string    `gorm:"type:varchar(100)" json:"name"`
	Login    string    `gorm:"type:varchar(100);unique" json:"login"`
	Password string    `gorm:"type:varchar(100)" json:"password"`
	Since    time.Time `json:"since"`
}

type UserDto struct {
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Login    string `gorm:"type:varchar(100);unique" json:"login"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}

// CreateUser creates new user
func FromDto(u UserDto) (User, error) {
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return User{}, err
	}

	return User{
		Name:     u.Name,
		Login:    u.Login,
		Password: hashedPassword,
		Since:    time.Now(),
	}, nil
}
