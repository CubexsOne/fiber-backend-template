package log

import (
	"log"
	"os"
)

var (
	Debug *log.Logger = log.New(
		os.Stdout,
		"🐞 DEBUG: ",
		log.Lshortfile,
	)
	Info *log.Logger = log.New(
		os.Stdout,
		"ℹ️ INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Warning *log.Logger = log.New(
		os.Stdout,
		"⚠️ WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Error *log.Logger = log.New(
		os.Stderr,
		"🔥 ERROR: ",
		log.Ldate|log.Ltime|log.Llongfile,
	)
)
