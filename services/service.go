package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"wife.storage/auth"
	"wife.storage/controllers"
	"wife.storage/database"
	"wife.storage/models"
)

// Service for data access
type Service struct {
}

// Run service
func (s *Service) Run() error {
	log.Print("Starting service ...")

	conf := &models.DatabaseConfig{}
	conf.Database = "public"
	conf.Host = "localhost"
	conf.Port = 5432
	conf.Provider = "postgres"
	conf.User = "postgres"
	conf.Password = "123456"

	port := 8081

	err := database.CreateDb(conf)
	if err != nil {
		log.Print("Failed to start service.")
		return err
	}

	router := mux.NewRouter()

	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	router.HandleFunc("/user", controllers.RegisterUser).Methods("POST")

	router.HandleFunc("/login", auth.Login).Methods("POST")

	http.ListenAndServe(":"+fmt.Sprint(port), router)

	log.Print("Service successfully started on port: ", port)

	return nil
}
