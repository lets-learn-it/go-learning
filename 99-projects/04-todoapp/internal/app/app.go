package app

import (
	"log"
	"net/http"
	"os"

	"avabodha.in/todoapp/internal/db"
	"avabodha.in/todoapp/internal/models"
	"github.com/go-playground/validator/v10"
)

var App Config

func NewApp(routes *http.Handler) *Config {
	// connect to the database
	db := db.InitDB()
	db.Ping()

	// Create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	App = Config{
		DB:            db,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Models:        models.New(db),
		Validator:     validator.New(),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	return &App
}
