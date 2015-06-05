package main

import (
	"fmt"
	"net/http"

	"github.com/arnoldcano/teaxdeax/infrastructure"
	"github.com/arnoldcano/teaxdeax/interfaces"
	"github.com/arnoldcano/teaxdeax/usecases"
	"github.com/gorilla/mux"
)

func main() {
	db := infrastructure.NewSqliteHandler("todos.sqlite")
	todoRepository := interfaces.NewTodoRepository(db)
	todoInteractor := usecases.NewTodoInteractor(todoRepository)
	todoHandler := interfaces.NewTodoHandler(todoInteractor)
	router := mux.NewRouter()
	todosRouter := router.Path("/todos").Subrouter()
	todosRouter.Methods("POST").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		todoHandler.Create(res, req)
	})
	todosRouter.Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		todoHandler.Index(res, req)
	})
	todoRouter := router.Path("/todo/{id}").Subrouter()
	todoRouter.Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		todoHandler.Show(res, req)
	})
	todoRouter.Methods("DELETE").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		todoHandler.Destroy(res, req)
	})
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
