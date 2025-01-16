package main

import (
	"log"
	"os"
	"url-shortener/db"
	"url-shortener/internal/user"
	"url-shortener/router"

	"github.com/joho/godotenv"
)

func main() {

	//? Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//? Connecting to Db
	dbConn, err := db.InitDb()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	defer dbConn.Close()

	//? Loading the server address
	port, isPresent := os.LookupEnv("PORT")
	if !isPresent {
		port = "8085"
	}

	addr := "0.0.0.0:" + port

	//? Creating the handlers
	userRep := user.NewUserRepository(dbConn)
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start(addr)

}
