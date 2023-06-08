package environment

import (
	"os"

	"github.com/cubexsone/fiber-backend-template/src/utils"
)

var (
	ENV         = utils.ToDefault(os.Getenv("ENV"), "DEV")
	PG_HOST     = utils.ToDefault(os.Getenv("PG_HOST"), "localhost")
	PG_PORT     = utils.ToDefault(os.Getenv("PG_PORT"), "5432")
	PG_USER     = utils.ToDefault(os.Getenv("PG_USER"), "prjctdb")
	PG_PASSWORD = utils.ToDefault(os.Getenv("PG_PASSWORD"), "prjctdb")
	PG_DATABASE = utils.ToDefault(os.Getenv("PG_DATABASE"), "prjctdb")
	PG_SCHEMA   = utils.ToDefault(os.Getenv("PG_SCHEMA"), "public")
	PG_SSL_MODE = utils.ToDefault(os.Getenv("PG_SSL_MODE"), "disable")
)
