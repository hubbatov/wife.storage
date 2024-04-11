package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wife/internal/errors"
)

type ServiceInterface interface {
	RegisterUser(dto UserDto) error
	Login(user string, password string) (bool, string)
	GetUsers() []User
}

type Endpoint struct {
	service ServiceInterface
}

func NewEndpoint(s ServiceInterface) *Endpoint {
	return &Endpoint{
		service: s,
	}
}

func (e *Endpoint) RegisterUser(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := io.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata UserDto
	err = json.Unmarshal(body, &userdata)
	errors.HandleError(errors.ConvertCustomError(err))

	err = e.service.RegisterUser(userdata)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Registered")
	}
}

func (e *Endpoint) Login(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := io.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata UserDto
	err = json.Unmarshal(body, &userdata)
	errors.HandleError(errors.ConvertCustomError(err))

	fl, token := e.service.Login(userdata.Login, userdata.Password)

	if fl {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "%s: %s", "Authorized", token)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Login or password incorrect")
	}
}

func (e *Endpoint) GetUsers(w http.ResponseWriter, req *http.Request) {
	fl := true
	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(e.service.GetUsers())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}
