package config

import (
	"log"

	"github.com/joho/godotenv"
)

// reads the .env file and sets the environment variables, if the file is not found, it will log an error and the program will default to injected production variables.
func ReadEnvfile() {
	//Read the .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Loading .env file failed, using production environment")
	}
	if err == nil {
		log.Println("Loaded .env file")
	}
}
