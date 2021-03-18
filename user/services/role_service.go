package services

import (
     "github.com/temaku/TodoApi/model"
	"github.com/temaku/TodoApi/user"
)

type RoleService struct {
	roleRepo user.RoleRepository
}

func NewRoleService(RoleRepo user.RoleRepository) *RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

func (rs *RoleService) Roles() ([]model.Role, []error) {

	rls, errs := rs.roleRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}
	return rls, errs

}

func (rs *RoleService) RoleByName(name string) (*model.Role, []error) {
	role, errs := rs.roleRepo.RoleByName(name)
	if len(errs) > 0 {
		return nil, errs
	}
	return role, errs
}

func (rs *RoleService) Role(id uint) (*model.Role, []error) {
	rl, errs := rs.roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

func (rs *RoleService) UpdateRole(role *model.Role) (*model.Role, []error) {
	rl, errs := rs.roleRepo.UpdateRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

func (rs *RoleService) DeleteRole(id uint) (*model.Role, []error) {

	rl, errs := rs.roleRepo.DeleteRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

func (rs *RoleService) StoreRole(role *model.Role) (*model.Role, []error) {

	rl, errs := rs.roleRepo.StoreRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}
