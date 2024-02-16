package main

import (
	"fmt"
	"log"
	"net/http"
	handler "todo/handlers"
	"todo/services"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("Starting server on the port 8080...")

	r.HandleFunc("/api/todo", handler.GetAllTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id}", handler.GetTodoById).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id}", handler.DeleteTodo).Methods(http.MethodDelete)
	r.HandleFunc("/api/upload", services.UploadFile).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
