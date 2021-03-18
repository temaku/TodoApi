package repository

import (
	"fmt"
	"github.com/temaku/TodoApi/model"
	"github.com/temaku/TodoApi/user"
	"github.com/temaku/TodoApi/utils"
	"github.com/jinzhu/gorm"
)

type UserGormRepo struct {
	conn *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

func (u *UserGormRepo) Users() ([]model.User, error) {
	var users []model.User
	errs := u.conn.Find(&users).Error
	if errs!=nil {
		return users, utils.ErrInternalServerError
	}
	return users, errs
}

func (u *UserGormRepo) User(id uint) (*model.User, error) {
	user1 := model.User{}
	errs := u.conn.First(&user1, id).Error
	if errs!=nil{
		return &user1, utils.ErrInternalServerError
	}
	return &user1, errs
}

func (u *UserGormRepo) StoreUser(user *model.User) (*model.User, error) {
	usr1:=user
	errs := u.conn.Create(&usr1).GetErrors()
	if len(errs)>0{
		println("Store User Gorm Exception")
		return nil, utils.ErrInternalServerError
	}
	return usr1, nil
}

func (u *UserGormRepo) UpdateUser(muser *model.User) (*model.User, error) {
	usr1:=muser
	errs := u.conn.Save(&usr1).Error
	if errs!=nil {
		return usr1, utils.ErrInternalServerError
	}
	return usr1, errs
}

func (u *UserGormRepo) DeleteUser(id uint) (*model.User, error) {
	user1, errs := u.User(id)
	if errs!=nil{
		return nil, utils.ErrInternalServerError
	}
	errs = u.conn.Delete(user1, id).Error
	if errs!=nil {
		return nil, utils.ErrInternalServerError
	}
	return user1, errs
}
func (u *UserGormRepo) UserByUserName(user model.User)(*model.User, error){
	user1 := model.User{}
	fmt.Println("gorm--- ",user)
	fmt.Println("email",user1.Email)
	errs := u.conn.Where("email = ?",user.Email).First(&user1).GetErrors()
	fmt.Println("gorm--- ",user1)
	if len(errs)>0{
		fmt.Println(errs)
		return &user1, utils.ErrInternalServerError
	}
	return &user1, nil
}
func (userRepo *UserGormRepo) PhoneExists(phone string) bool {
	user := model.User{}
	errs := userRepo.conn.Find(&user, "phone=?", phone).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

func (userRepo *UserGormRepo) EmailExists(email string) bool {
	user := model.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

func (userRepo *UserGormRepo) UserRoles(user *model.User) ([]model.Role, []error) {
	userRoles := []model.Role{}
	errs := userRepo.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}





