package repository

import (
	"github.com/temaku/TodoApi/model"
	"github.com/jinzhu/gorm"
)

type RoleGormRepo struct {
	conn *gorm.DB
}

func NewRoleGormRepo(db *gorm.DB) *RoleGormRepo {
	return &RoleGormRepo{conn: db}
}

func (roleRepo *RoleGormRepo) Roles() ([]model.Role, []error) {
	roles := []model.Role{}
	errs := roleRepo.conn.Find(&roles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return roles, errs
}

func (roleRepo *RoleGormRepo) Role(id uint) (*model.Role, []error) {
	role := model.Role{}
	errs := roleRepo.conn.First(&role, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}

func (roleRepo *RoleGormRepo) RoleByName(name string) (*model.Role, []error) {
	role := model.Role{}
	errs := roleRepo.conn.Find(&role, "name=?", name).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}

func (roleRepo *RoleGormRepo) UpdateRole(role *model.Role) (*model.Role, []error) {
	r := role
	errs := roleRepo.conn.Save(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

func (roleRepo *RoleGormRepo) DeleteRole(id uint) (*model.Role, []error) {
	r, errs := roleRepo.Role(uint(int(id)))
	if len(errs) > 0 {
		return nil, errs
	}
	errs = roleRepo.conn.Delete(r, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}

func (roleRepo *RoleGormRepo) StoreRole(role *model.Role) (*model.Role, []error) {
	r := role
	errs := roleRepo.conn.Create(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}
