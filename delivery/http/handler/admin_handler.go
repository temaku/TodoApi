package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/temaku/TodoApi/todo"
	"github.com/temaku/TodoApi/delivery/http/response"
	"github.com/temaku/TodoApi/model"

	//"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	tosr        todo.TodoService
	csrfSignKey []byte
}

func NewAdminHandler(td todo.TodoService) *AdminHandler {

	return &AdminHandler{tosr: td }

}

//Todo
func (ah *AdminHandler) AdminTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, _ := strconv.Atoi(id)

	cinema, errs := ah.tosr.Todo(uint(idd))
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Cinema"))
	}
	responses.JSON(w, http.StatusOK, cinema)
}

func (ah *AdminHandler) AdminTodos(w http.ResponseWriter, r *http.Request) {
	cinemas, errs := ah.tosr.Todos()
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Cinemas"))
	}
	responses.JSON(w, http.StatusOK, cinemas)
}

func (ah *AdminHandler) AdminDeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, _ := strconv.Atoi(id)

	cinema, errs := ah.tosr.DeleteTodo(uint(idd))
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Delete Cinema"))
	}
	responses.JSON(w, http.StatusNoContent, cinema)

}
func (ah *AdminHandler) AdminTodoUpdateList(w http.ResponseWriter, r *http.Request){
	var c model.Todo
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}

	idd, err := strconv.Atoi(id)
	cin, errs := ah.tosr.Todo(uint(idd))

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant fetch cinema"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	cin, errs = ah.tosr.UpdateTodo(&c)

	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant update cinema"))
		return
	}
	responses.JSON(w, http.StatusOK, cin)
	return
}

func (ah *AdminHandler) AdminTodoNew(w http.ResponseWriter, r *http.Request){
	var c model.Todo
	err := json.NewDecoder(r.Body).Decode(&c)
	if err!=nil{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("todo decoding failed"))
		return
	}
	if ah.tosr.TodoExists(c.Title){
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Todo already exits"))
		return
		//json.NewEncoder(w).Encode(err)
	}

	newCinema, errs := ah.tosr.StoreTodo(&c)
	if len(errs) > 0 {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Adding new Cinema Failed"))
		return
	}

	responses.JSON(w, http.StatusOK, newCinema)
	//json.NewEncoder(w).Encode(newCinema)
}



