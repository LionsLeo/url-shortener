package main

import (
	"log"
	"os"
	"url-shortener/router"

	"github.com/joho/godotenv"
)

func main() {

	//? Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//? Loading the server address
	port, isPresent := os.LookupEnv("PORT")
	if !isPresent {
		port = "8085"
	}

	addr := "0.0.0.0:" + port

	router.InitRouter()
	router.Start(addr)

}
