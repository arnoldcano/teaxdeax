package interfaces

import (
	"encoding/json"
	"net/http"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/arnoldcano/teaxdeax/domain"
	"github.com/arnoldcano/teaxdeax/usecases"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	interactor *usecases.TodoInteractor
}

func NewTodoHandler(interactor *usecases.TodoInteractor) *TodoHandler {
	return &TodoHandler{
		interactor: interactor,
	}
}

func (handler *TodoHandler) Create(res http.ResponseWriter, req *http.Request) {
	id := uuid.New()
	note := req.FormValue("note")
	now := time.Now()
	todo := domain.NewTodo(id, note, now, now)
	err := handler.interactor.Create(todo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *TodoHandler) Index(res http.ResponseWriter, req *http.Request) {
	todos, err := handler.interactor.FindAll()
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	js, err := json.Marshal(todos)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func (handler *TodoHandler) Show(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	todo, err := handler.interactor.FindById(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	js, err := json.Marshal(todo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func (handler *TodoHandler) Destroy(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := handler.interactor.DeleteById(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
}
