package interfaces

import (
	"net/http"

	"code.google.com/p/go-uuid/uuid"

	"github.com/arnoldcano/teaxdeax/domain"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type TodosInteractor interface {
	Create(*domain.Todo) error
	FindAll() ([]*domain.Todo, error)
	FindById(string) (*domain.Todo, error)
	Update(*domain.Todo) error
	DeleteById(string) error
}

type TodosHandler struct {
	interactor TodosInteractor
}

func NewTodosHandler(interactor TodosInteractor) *TodosHandler {
	return &TodosHandler{
		interactor: interactor,
	}
}

func (h *TodosHandler) Create(res http.ResponseWriter, req *http.Request) {
	id := uuid.New()
	note := req.FormValue("note")
	todo := domain.NewTodo(id, note)
	if err := h.interactor.Create(todo); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodosHandler) Index(res http.ResponseWriter, req *http.Request) {
	todos, err := h.interactor.FindAll()
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	render := render.New()
	render.JSON(res, http.StatusOK, todos)
}

func (h *TodosHandler) Show(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	todo, err := h.interactor.FindById(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	render := render.New()
	render.JSON(res, http.StatusOK, todo)
}

func (h *TodosHandler) Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	note := req.FormValue("note")
	todo := domain.NewTodo(id, note)
	if err := h.interactor.Update(todo); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodosHandler) Destroy(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	if err := h.interactor.DeleteById(id); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
