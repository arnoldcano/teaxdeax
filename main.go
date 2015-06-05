package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arnoldcano/teaxdeax/infrastructure"
	"github.com/arnoldcano/teaxdeax/interfaces"
	"github.com/arnoldcano/teaxdeax/usecases"
	"github.com/gorilla/mux"
)

func main() {
	db := infrastructure.NewSqliteHandler(os.Getenv("FILE"))
	repo := interfaces.NewTodosRepository(db)
	interactor := usecases.NewTodosInteractor(repo)
	handler := interfaces.NewTodosHandler(interactor)
	mux := mux.NewRouter()
	todos := mux.Path("/todos").Subrouter()
	todos.Methods("POST").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.Create(res, req)
	})
	todos.Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.Index(res, req)
	})
	todo := mux.Path("/todos/{id}").Subrouter()
	todo.Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.Show(res, req)
	})
	todo.Methods("PUT").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.Update(res, req)
	})
	todo.Methods("DELETE").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler.Destroy(res, req)
	})
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
