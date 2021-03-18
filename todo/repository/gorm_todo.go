package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/temaku/TodoApi/model"
)

type TodoGormRepo struct {
	conn *gorm.DB
}


func NewTodoGormRepo(db *gorm.DB) *TodoGormRepo {
	return &TodoGormRepo{conn: db}
}

func (todoRepo *TodoGormRepo) Todos() ([]model.Todo, []error) {
	var clubs []model.Todo
	errs := todoRepo.conn.Find(&clubs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return clubs, errs
}

func (todoRepo *TodoGormRepo) Todo(id uint) (*model.Todo, []error) {
	club := model.Todo{}
	errs := todoRepo.conn.First(&club, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &club, errs
}

func (todoRepo *TodoGormRepo) UpdateTodo(todo *model.Todo) (*model.Todo, []error) {
	errs := todoRepo.conn.Save(todo).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return todo, errs
}

func (todoRepo *TodoGormRepo) DeleteTodo(id uint) (*model.Todo, []error) {
	club, errs := todoRepo.Todo(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = todoRepo.conn.Delete(club, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return club, errs
}

func (todoRepo *TodoGormRepo) StoreTodo(todo *model.Todo) (*model.Todo, []error) {
	errs := todoRepo.conn.Create(todo).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return todo, errs
}
func (todoRepo *TodoGormRepo) TodoExists(todoname string) bool {
	eve := model.Todo{}
	errs := todoRepo.conn.Find(&eve, "title=?", todoname).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}






