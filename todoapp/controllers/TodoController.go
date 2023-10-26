package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/emanet/gotest/todoapp/database"
	"github.com/emanet/gotest/todoapp/models"
)

func TodoList(w http.ResponseWriter, r *http.Request) {
	view, err := template.ParseFiles("views/list.html")
	if err != nil {
		log.Fatal("Template not found ")
	}

	db := database.Init()
	var todos []models.Todo

	query := db.Find(&todos)

	defer query.Close()

	view.Execute(w, todos)
}
func TodoAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		TodoForm(w, r)
	case "POST":
		ProcessTodoForm(w, r)
	}
}

func TodoForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/add.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func ProcessTodoForm(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		http.Redirect(w, r, "/", 302)
	}
	db := database.Init()
	todo := models.Todo{
		Title:     title,
		Completed: false,
	}
	db.Create(&todo)
	http.Redirect(w, r, "/", 302)
}
