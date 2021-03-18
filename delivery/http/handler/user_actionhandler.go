package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/temaku/TodoApi/todo"
	"github.com/temaku/TodoApi/delivery/http/response"
	"github.com/temaku/TodoApi/model"
	"github.com/temaku/TodoApi/rtoken"
	"github.com/temaku/TodoApi/user"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
)

type UserActionHandler struct {

	torv   todo.TodoService
	usrv   user.UsersService
	tsrv   rtoken.Service

}

func NewUserActionHandler(cs todo.TodoService, u user.UsersService, t rtoken.Service) *UserActionHandler {
	return &UserActionHandler{torv: cs,tsrv:t, usrv: u}
}
func (uah *UserActionHandler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	var u model.User
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)
	//log.Println(idd)
	usr, errs := uah.usrv.User(uint(idd))

	curid := usr.Id
	pass := usr.Password
	rol := usr.RoleID

	log.Println(curid)
	log.Println(pass)
	log.Println(rol)
	log.Println(usr.FullName)

	if errs != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant fetch user"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	passnew, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Password Encryption  failed"))
		return
	}

	u.Password = string(passnew)
	u.RoleID = rol


	usr, errs = uah.usrv.UpdateUser(&u)

	log.Println(usr.Id)
	log.Println(usr.Password)
	log.Println(usr.RoleID)
	log.Println(usr.FullName)

	if errs != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant update user"))
		return
	}
	responses.JSON(w, http.StatusOK, &usr)
	return
}

func (uah *UserActionHandler) UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not given"))
		return
	}
	idd, _ := strconv.Atoi(id)
	_, errs := uah.usrv.DeleteUser(uint(idd))
	if errs != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Couldn't delete user"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

func (uah *UserActionHandler) Todos(w http.ResponseWriter, r *http.Request){
	//var mov []model.Movie
	nowshowingmovies, errs := uah.torv.Todos()
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Movies"))
	}
	responses.JSON(w, http.StatusOK, nowshowingmovies)
}

func (uah *UserActionHandler) SingleTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, exists := params["id"]
	if !exists {
		responses.ERROR(w, http.StatusBadRequest, errors.New("id not provided"))
		return
	}
	idd, err := strconv.Atoi(id)

	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Can't get movie ID"))
	}

	curMovie, errs := uah.torv.Todo(uint(idd))
	if len(errs) > 0{
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Cant Fetch Movies"))
	}
	responses.JSON(w, http.StatusOK, curMovie)
}


