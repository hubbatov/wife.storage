package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"wife.storage/auth"
	"wife.storage/database"
	"wife.storage/errors"
	"wife.storage/rest"

	"strings"

	"fmt"
)

// GetUsers returns json with all users in system
func GetUsers(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, _ := auth.CheckAuthorization(accessToken)

	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(database.DBManager.Users())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

// GetUser returns json with authorized user data
func GetUser(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, userID := auth.CheckAuthorization(accessToken)

	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(auth.GetUserByID(userID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

// RegisterUser creates new user in system
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata rest.RESTUser
	err = json.Unmarshal(body, &userdata)
	errors.HandleError(errors.ConvertCustomError(err))

	eArray := database.DBManager.CreateUser(userdata)

	if len(eArray) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", eArray)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Registered")
	}
}
