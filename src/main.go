package main

import (
	"os"

	"github.com/cubexsone/fiber-backend-template/src/database"
	"github.com/cubexsone/fiber-backend-template/src/server"
	"github.com/cubexsone/fiber-backend-template/src/utils/log"
)

func main() {
	log.Info.Println("Server is starting...")
	database.Connect()
	log.Info.Println("ENV: " + os.Getenv("ENV"))
	server.Server()
}
