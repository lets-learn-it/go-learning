package handler

import (
	"net/http"
	"strconv"
	"time"

	"avabodha.in/todoapp/internal/app"
	"avabodha.in/todoapp/internal/models"
	"github.com/go-chi/chi/v5"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	var todo *models.Todo

	ids := chi.URLParam(r, "id")

	app.App.InfoLog.Printf("Got todo id: %s.", ids)

	id, err := strconv.Atoi(ids)

	if err != nil {
		app.App.ErrorLog.Printf("Unable to convert %s to integer.", ids)
	}

	todo, err = app.App.Models.Todo.GetOne(id)

	if err != nil {
		app.App.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	todo.ToJSON(w)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	todo.FromJSON(r.Body)

	err := app.App.Validator.Struct(todo)

	if err != nil {
		app.App.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.CreatedAt = time.Now()
	todo.ModifiedAt = time.Now()

	newId, err := app.App.Models.Todo.Insert(todo)

	if err != nil {
		app.App.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todo.Id = int64(newId)

	w.Header().Set("content-type", "application/json")
	todo.ToJSON(w)
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	todo.FromJSON(r.Body)

	err := app.App.Validator.Struct(todo)

	if err != nil {
		app.App.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ModifiedAt = time.Now()

	ids := chi.URLParam(r, "id")

	app.App.InfoLog.Printf("Got todo id: %s for updation.", ids)

	id, err := strconv.Atoi(ids)

	if err != nil {
		app.App.ErrorLog.Printf("Unable to convert %s to integer.", ids)
	}

	err = app.App.Models.Todo.Update(id, todo)

	if err != nil {
		app.App.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ids := chi.URLParam(r, "id")

	app.App.InfoLog.Printf("Got todo id: %s for deletion.", ids)

	id, err := strconv.Atoi(ids)

	if err != nil {
		app.App.ErrorLog.Printf("Unable to convert %s to integer.", ids)
	}

	err = app.App.Models.Todo.DeleteByID(id)

	if err != nil {
		app.App.ErrorLog.Println(err)
	}

}
