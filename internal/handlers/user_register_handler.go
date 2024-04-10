package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"wife/internal/api"
	"wife/internal/database"
	"wife/internal/errors"

	"strings"

	"fmt"
)

// RegisterUser creates new user in system
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := io.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata api.User
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
