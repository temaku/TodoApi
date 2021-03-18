package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/temaku/TodoApi/delivery/http/handler"
	"github.com/temaku/TodoApi/model"
	"github.com/temaku/TodoApi/rtoken"

	usrepo "github.com/temaku/TodoApi/user/repository"
	userv "github.com/temaku/TodoApi/user/services"

	mrepo "github.com/temaku/TodoApi/todo/repository"
	mserv "github.com/temaku/TodoApi/todo/service"
)

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:7409@localhost/todoItems?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Todo{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})

	db.AutoMigrate(&model.Role{ID: 1, Name: "USER"})
	db.AutoMigrate(&model.Role{ID: 2, Name: "ADMIN"})

	token := rtoken.Service{}


	MovieRepo := mrepo.NewTodoGormRepo(db)
	Moviesr := mserv.NewTodoService(MovieRepo)

	UserRepo := usrepo.NewUserGormRepo(db)
	usersr := userv.NewUserService(UserRepo)

	roleRepo := usrepo.NewRoleGormRepo(db)
	rolesr := userv.NewRoleService(roleRepo)

	//ch := handler.NewCinemaHandler(Cinemasr)
	//sch := handler.NewScheduleHandler(Schedulesr)
	//boh := handler.NewBookingHandler(Bookingsr)

	uh := handler.NewUserHandler(usersr, rolesr, token)
	uah := handler.NewUserActionHandler( Moviesr, usersr, token)
	ah := handler.NewAdminHandler(Moviesr)

	router := mux.NewRouter()



	router.HandleFunc("/todos", ah.AdminTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", ah.AdminTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", ah.AdminTodoUpdateList).Methods("PUT")
	router.HandleFunc("/todo/{id}", ah.AdminDeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo", ah.AdminTodoNew).Methods("POST")



	router.HandleFunc("/login", uh.Login).Methods("POST")
	router.HandleFunc("/signup", uh.SignUp).Methods("POST")
	router.HandleFunc("/logout", uh.Authenticated(uh.Logout)).Methods("POST")
	router.HandleFunc("/user/{id}", uh.Authenticated(uah.UserUpdate)).Methods("PUT")
	router.HandleFunc("/user/{id}", uh.Authenticated(uah.UserDelete)).Methods("DELETE")


	http.ListenAndServe(":8181", router)

}
