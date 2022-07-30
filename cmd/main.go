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
	config.Env()
	//Database adapter - //internal/db
	redisdb, err := repository.NewRedisAdapter()
	fmt.Println("Redis adapter:", redisdb)
	if err != nil {
		fmt.Println(err)
	}

	mongodb, err := repository.NewMongoAdapter()
	fmt.Println("DB adapter initialized: ", mongodb)
	if err != nil {
		fmt.Println(err)
	}

	filedb := filerepository.NewFileAdapter()
	fmt.Println("File adapter initialized: ", filedb)

	//Application //internal/application/api
	applicationAPI := api.NewApplication(mongodb, mongodb, filedb, redisdb)
	fmt.Println("API adapter initialized: ", applicationAPI)

	//serverAdapter - //internal/server
	serverAdapter := server.NewServerAdapter(applicationAPI, applicationAPI)
	fmt.Println("Server adapter initialized: ", serverAdapter)
	server.Router(serverAdapter)
}
