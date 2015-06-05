package main

import (
	"fmt"
	"net/http"

	"github.com/arnoldcano/teaxdeax/infrastructure"
	"github.com/arnoldcano/teaxdeax/interfaces"
	"github.com/arnoldcano/teaxdeax/usecases"
	"github.com/gorilla/mux"
)

var todoHandler *interfaces.TodoHandler

func main() {
	db := infrastructure.NewSqliteHandler("todos.sqlite")
	todoRepository := interfaces.NewTodoRepository(db)
	todoInteractor := usecases.NewTodoInteractor(todoRepository)
	todoHandler = interfaces.NewTodoHandler(todoInteractor)
	router := mux.NewRouter()
	todos := router.Path("/todos").Subrouter()
	todos.Methods("GET").HandlerFunc(TodosIndexHandler)
	todos.Methods("POST").HandlerFunc(TodosCreateHandler)
	todo := router.Path("/todos/{id}").Subrouter()
	todo.Methods("GET").HandlerFunc(TodoShowHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}

func TodosCreateHandler(res http.ResponseWriter, req *http.Request) {
	todoHandler.Create(res, req)
}

func TodosIndexHandler(res http.ResponseWriter, req *http.Request) {
	todoHandler.FindAll(res, req)
}

func TodoShowHandler(res http.ResponseWriter, req *http.Request) {
	todoHandler.FindById(res, req)
}
