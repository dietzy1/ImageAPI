package main

import (
	"fmt"

	//Adapters
	adapter "github.com/dietzy1/imageAPI/internal/adaptors"
	"github.com/dietzy1/imageAPI/internal/config"

	"github.com/dietzy1/imageAPI/internal/application/api"
	//Application
)

func main() {
	config.Env()
	//Database adapter - //internal/db
	mongodb, err := adapter.NewDbAdapter()
	fmt.Println("DB adapter initialized: ", mongodb)
	if err != nil {
		fmt.Println(err)
	}

	filedb := adapter.NewFileAdapter()
	fmt.Println("File adapter initialized: ", filedb)

	//Application //internal/application/api
	applicationAPI := api.NewApplication(mongodb, filedb)
	fmt.Println("API adapter initialized: ", applicationAPI)

	//serverAdapter - //internal/server
	serverAdapter := adapter.NewServerAdapter(applicationAPI)
	fmt.Println("Server adapter initialized: ", serverAdapter)
	adapter.Router(serverAdapter)

}
