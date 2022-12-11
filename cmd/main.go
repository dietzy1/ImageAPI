package main

import (
	"fmt"

	//Adapters
	"github.com/dietzy1/imageAPI/internal/adaptors/filerepository"
	"github.com/dietzy1/imageAPI/internal/adaptors/repository"

	//Adapters
	"github.com/dietzy1/imageAPI/internal/adaptors/server"

	//config
	"github.com/dietzy1/imageAPI/internal/config"

	//Application
	"github.com/dietzy1/imageAPI/internal/application/api"
)

func main() {
	config.ReadEnvfile()

	//Database adapter
	redisdb, err := repository.NewRedisAdapter()
	fmt.Println("Redis adapter:", redisdb)
	if err != nil {
		fmt.Println(err)
	}

	//database adapter
	mongodb, err := repository.NewMongoAdapter()
	fmt.Println("DB adapter initialized: ", mongodb)
	if err != nil {
		fmt.Println(err)
	}

	//Cdn adapter
	cdn, err := filerepository.NewImageKitClientAdapter()
	fmt.Println("File adapter initialized: ", cdn)
	if err != nil {
		fmt.Println(err)
	}

	//Application //internal/application/api
	applicationAPI := api.NewApplication(mongodb, mongodb, mongodb, mongodb, redisdb, cdn)
	fmt.Println("API adapter initialized: ", applicationAPI)

	//serverAdapter - //internal/server
	serverAdapter := server.NewServerAdapter(applicationAPI, applicationAPI, applicationAPI, applicationAPI)
	fmt.Println("Server adapter initialized: ", serverAdapter)
	server.Router(serverAdapter)
}

//TODO:

//I want to reconfigure the database
//Every user should be able to claim ownershio of a pepe
//users should be able to upload their own pepe
//User struct should contain a list of pepe ids
//pepe should contain the owners id
//pepe should contain the owners username
