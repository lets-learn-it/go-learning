package app

import (
	"database/sql"

	"log"

	"avabodha.in/todoapp/internal/models"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Models        models.Models
	Validator     *validator.Validate
	ErrorChan     chan error
	ErrorChanDone chan bool
}
