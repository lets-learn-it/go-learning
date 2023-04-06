package main

import (
	"log"
	"net/http"
	"os"

	"avabodha.in/go-fiber-postgres/models"
	"avabodha.in/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed. " + err.Error()})
		return err
	}

	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book. " + err.Error()})
		return err
	}

	context.Status(http.StatusCreated).JSON(
		&fiber.Map{"message": "book created."})

	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "provide parameter id"})
		return nil
	}

	err := r.DB.Delete(bookModel, id).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "got error while removing book. " + err.Error()})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book deleted."})

	return nil
}

func (r *Repository) GetAllBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "got error while getting books. " + err.Error()})
		return err
	}

	context.Status(http.StatusOK).JSON(bookModels)
	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	id := context.Params("id")

	bookModel := &models.Books{}

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "provide parameter id"})
		return nil
	}

	log.Println("The Id is ", id)
	err := r.DB.Where("id = ?", id).First(bookModel).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "got error while getting book with id: " + id + ". " + err.Error()})
		return err
	}

	context.Status(http.StatusOK).JSON(bookModel)
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/book", r.CreateBook)
	api.Delete("/book/:id", r.DeleteBook)
	api.Get("/book/:id", r.GetBookById)
	api.Get("/book", r.GetAllBooks)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)

	app.Listen(":8080")
}
