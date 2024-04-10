package services

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"wife/internal/database"
	"wife/internal/handlers"
)

type AuthService struct {
}

// Run service
func (s *AuthService) Run() error {
	log.Print("Starting auth service ...")

	conf := &database.DatabaseConfig{}
	conf.Database = os.Getenv("WIFEDB")
	conf.Host = os.Getenv("WIFEHOST")
	conf.Port, _ = strconv.Atoi(os.Getenv("WIFEPORT"))
	conf.Provider = os.Getenv("WIFEPROVIDER")
	conf.User = os.Getenv("WIFEEUSER")
	conf.Password = os.Getenv("WIFEPASSWORD")

	port := 8081

	err := database.CreateDb(conf)
	if err != nil {
		log.Print("Failed to start service.")
		return err
	}

	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/user", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	http.ListenAndServe(":"+fmt.Sprint(port), router)

	log.Print("Service successfully started on port: ", port)

	return nil
}
