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
	fmt.Println("Starting server...")
	config.Env()
	//Database adapter - //internal/db
	mongodb, err := adapter.NewDbAdapter()
	if err != nil {
	}
	filedb := adapter.NewFileAdapter()

	fmt.Println(mongodb, "this is db adapter")

	//Application //internal/application/api
	applicationAPI := api.NewApplication(mongodb, filedb)
	fmt.Println(applicationAPI, "This is application API")

	//serverAdapter - //internal/server
	serverAdapter := adapter.NewServerAdapter(applicationAPI)
	fmt.Println(serverAdapter, "This is serveradapter")
	adapter.Router(serverAdapter)

	//Iniate frontend

}

//Potentially use filepaths in the db and then read in the pepes from the disk
