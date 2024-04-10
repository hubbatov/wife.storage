package database

import (
	"time"
	"wife/internal/api"
)

// User represents typical user of service
type User struct {
	ID             int       `gorm:"primary_key" json:"id"`
	Name           string    `gorm:"type:varchar(100)" json:"name"`
	Login          string    `gorm:"type:varchar(100);unique" json:"login"`
	Password       string    `gorm:"type:varchar(100)" json:"password"`
	Since          time.Time `json:"since"`
	HashedPassword []byte    `json:"-"`
}

// CreateUser creates new user
func CreateUser(userdata api.User) User {
	a := User{}
	a.Name = userdata.Name
	a.Login = userdata.Login
	a.Password = userdata.Password
	a.Since = time.Now()
	return a
}

// TableName for Users
func (User) TableName() string {
	return "public.users"
}
