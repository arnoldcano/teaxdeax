package interfaces

import (
	"net/http"

	"code.google.com/p/go-uuid/uuid"

	"github.com/arnoldcano/teaxdeax/domain"
	"github.com/arnoldcano/teaxdeax/usecases"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type TodoHandler struct {
	interactor *usecases.TodoInteractor
}

func NewTodoHandler(interactor *usecases.TodoInteractor) *TodoHandler {
	return &TodoHandler{
		interactor: interactor,
	}
}

func (h *TodoHandler) Create(res http.ResponseWriter, req *http.Request) {
	id := uuid.New()
	note := req.FormValue("note")
	todo := domain.NewTodo(id, note)
	err := h.interactor.Create(todo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodoHandler) Index(res http.ResponseWriter, req *http.Request) {
	todos, err := h.interactor.FindAll()
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	render := render.New()
	render.JSON(res, http.StatusOK, todos)
}

func (h *TodoHandler) Show(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	todo, err := h.interactor.FindById(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	render := render.New()
	render.JSON(res, http.StatusOK, todo)
}

func (h *TodoHandler) Update(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	note := req.FormValue("note")
	todo := domain.NewTodo(id, note)
	err := h.interactor.Update(todo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TodoHandler) Destroy(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := h.interactor.DeleteById(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
}
