package main

import (
	"wife.storage/services"
)

func main() {
	s := new(services.Service)
	err := s.Run()
	if err != nil {
		return
	}
}
