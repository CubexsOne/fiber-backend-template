package database

import (
	"github.com/cubexsone/fiber-backend-template/src/environment"
	"github.com/cubexsone/fiber-backend-template/src/utils/log"
)

var (
	host     = environment.PG_HOST
	port     = environment.PG_PORT
	user     = environment.PG_USER
	password = environment.PG_PASSWORD
	database = environment.PG_DATABASE
	schema   = environment.PG_SCHEMA
	sslMode  = environment.PG_SSL_MODE
)

func Connect() {
	database := New(host, port, user, password, database, schema, sslMode)
	err := database.Connect()
	if err != nil {
		log.Error.Println("Connect to database failed:", err)
	}
}
