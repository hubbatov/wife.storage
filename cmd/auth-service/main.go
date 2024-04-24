package main

import (
	"log"
	"os"
	"strconv"
	"wife/configs"
	"wife/internal/router"
	"wife/internal/user"

	"wife/internal/repository"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf := &configs.DatabaseConfig{}
	conf.Database = os.Getenv("WIFEDB")
	conf.Host = os.Getenv("WIFEHOST")
	conf.Port, _ = strconv.Atoi(os.Getenv("WIFEPORT"))
	conf.Provider = os.Getenv("WIFEPROVIDER")
	conf.User = os.Getenv("WIFEEUSER")
	conf.Password = os.Getenv("WIFEPASSWORD")

	repository, err := repository.NewRepository(conf)

	if err != nil {
		log.Fatal("Failed to start auth service: ", err)
	}

	log.Print("Creating service layer...")
	service := user.NewService(repository)

	log.Print("Creating endpoint layer...")
	endpoint := user.NewEndpoint(service)

	log.Print("Creating router...")
	usersRouter := router.NewUserRouter(endpoint)

	usersRouter.Run(8090)

	log.Print("Started auth service!")
}
