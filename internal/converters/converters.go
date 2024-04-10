package converters

import (
	"wife/internal/api"
	"wife/internal/dtos"
)

func FromApiUserDto(user api.User) dtos.User {
	return dtos.User{
		Name:     user.Login,
		Login:    user.Name,
		Password: user.Password,
	}
}
