package services

import (
	"fmt"
	"github.com/temaku/TodoApi/model"
	"github.com/temaku/TodoApi/user"
)

type UserServices struct {
	userRepo user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserServices {
	return &UserServices{userRepo: userRepository}
}
func (u *UserServices) Users() ([]model.User, error) {
	users, errs := u.userRepo.Users()
	if errs!=nil{
		return users, errs
	}
	return users, nil
}

func (u *UserServices) User(id uint) (*model.User, error) {
	users, errs := u.userRepo.User(id)
	if errs!=nil{
		return users, errs
	}
	return users, nil
}

func (u *UserServices) UpdateUser(user *model.User) (*model.User, error) {
	user1, errs := u.userRepo.UpdateUser(user)
	if errs!=nil {
		return user1, errs
	}
	return user1, nil
}

func (u *UserServices) DeleteUser(id uint) (*model.User, error) {
	user1, errs := u.userRepo.DeleteUser(id)
	if errs!=nil {
		fmt.Println("Delete Room Services")
		return user1, errs
	}
	return user1, nil
}

func (u *UserServices) StoreUser(user *model.User) (*model.User, error) {
	us:=user
	user1, errs := u.userRepo.StoreUser(us)
	if errs!=nil {
		return  nil,errs
	}
	return user1, nil
}
func (u *UserServices) UserByUserName(user model.User)(*model.User, error){
	users, errs := u.userRepo.UserByUserName(user)
	if errs!=nil{
		return users, errs
	}
	fmt.Println("services--- ",*users)
	return users, nil
}


func (us *UserServices) PhoneExists(phone string) bool {
	exists := us.userRepo.PhoneExists(phone)
	return exists
}

func (us *UserServices) EmailExists(email string) bool {
	exists := us.userRepo.EmailExists(email)
	return exists
}

func (us *UserServices) UserRoles(user *model.User) ([]model.Role, []error) {
	userRoles, errs := us.userRepo.UserRoles(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}
