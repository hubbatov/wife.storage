package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"wife.storage/auth"
	"wife.storage/controllers"
	"wife.storage/database"
)

// Service for data access
type Service struct {
}

// Run service
func (s *Service) Run() {
	log.Print("Starting service ...")

	database.CreateDb()

	router := mux.NewRouter()

	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.RegisterUser).Methods("POST")

	router.HandleFunc("/login", auth.Login).Methods("POST")

	http.ListenAndServe(":8081", router)
}
