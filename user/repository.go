package user

import "github.com/temaku/TodoApi/model"

type UserRepository interface {
	Users() ([]model.User, error)
	User(id uint) (*model.User, error)
	StoreUser(user *model.User) (*model.User, error)
	UpdateUser(order *model.User) (*model.User, error)
	DeleteUser(id uint) (*model.User, error)
	UserByUserName(user model.User)(*model.User, error)
	PhoneExists(phone string) bool
	EmailExists(email string) bool
	UserRoles(*model.User) ([]model.Role, []error)
}

type RoleRepository interface {
	Roles() ([]model.Role, []error)
	Role(id uint) (*model.Role, []error)
	RoleByName(name string) (*model.Role, []error)
	UpdateRole(role *model.Role) (*model.Role, []error)
	DeleteRole(id uint) (*model.Role, []error)
	StoreRole(role *model.Role) (*model.Role, []error)
}


