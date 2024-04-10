package handlers

import (
	"encoding/json"
	"net/http"

	"wife/internal/database"

	"fmt"
)

// GetUsers returns json with all users in system
func GetUsers(w http.ResponseWriter, req *http.Request) {
	//accessToken := req.Header.Get("Authorization")

	//fl, _ := auth.CheckAuthorization(accessToken)
	fl := true
	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(database.DBManager.Users())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}
