package user

import (
	"fmt"
)

type RepositoryInterface interface {
	CreateUser(dto UserDto) []error
	Users() []User
	Exists(login string, password string) bool
}

type Service struct {
	repository RepositoryInterface
}

func NewService(r RepositoryInterface) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) RegisterUser(dto UserDto) error {

	eArray := s.repository.CreateUser(dto)

	if len(eArray) > 0 {
		return fmt.Errorf("%s", eArray)
	}

	return nil
}

func (s *Service) Login(login string, password string) (bool, string) {
	token := "not implemented"
	loggedIn := s.repository.Exists(login, password)
	return loggedIn, token
}

func (s *Service) GetUsers() []User {
	return s.repository.Users()
}
