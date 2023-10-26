package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.TodoList)
	r.HandleFunc("/add", controllers.TodoAdd)
	r.HandleFunc("/delete/{id}", controllers.TodoDelete)
	http.Handle("/", r)

	fmt.Println("Server started")
	http.ListenAndServe(":3000", nil)
}
