package main

import (
	"wife/internal/services"
)

func main() {
	s := new(services.AuthService)
	err := s.Run()
	if err != nil {
		return
	}
}
