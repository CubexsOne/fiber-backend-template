package database

import (
	"fmt"

	"github.com/cubexsone/fiber-backend-template/src/utils/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var instance *sqlx.DB

type Database struct {
	config *config
}

type config struct {
	host     string
	port     string
	user     string
	password string
	database string
	schema   string
	sslMode  string
}

func New(
	host string,
	port string,
	user string,
	password string,
	database string,
	schema string,
	sslMode string,
) *Database {
	return &Database{
		config: &config{
			user:     user,
			password: password,
			port:     port,
			database: database,
			host:     host,
			schema:   schema,
			sslMode:  sslMode,
		},
	}
}

func (database *Database) Connect() error {
	config := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s  search_path=%s sslmode=%s",
		database.config.host,
		database.config.port,
		database.config.user,
		database.config.password,
		database.config.database,
		database.config.schema,
		database.config.sslMode,
	)

	openInstance, openErr := sqlx.Open("postgres", config)
	if openErr != nil {
		return openErr
	}

	conErr := openInstance.Ping()
	instance = openInstance

	return conErr
}

func GetInstance() *sqlx.DB {
	if instance != nil {
		log.Error.Println("No database connection established!")
	}

	return instance
}
