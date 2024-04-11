package main

import (
	"log"
	"os"
	"strconv"
	"wife/configs"
	"wife/internal/router"
	"wife/internal/user"

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

	repository, err := user.NewRepository(conf)

	if err != nil {
		log.Fatal("Failed to start service: ", err)
	}

	service := user.NewService(repository)
	endpoint := user.NewEndpoint(service)

	usersRouter := router.NewUserRouter(endpoint)
	err = usersRouter.Run(8081)
	if err != nil {
		log.Fatal("Failed to start service: ", err)
	}
}
