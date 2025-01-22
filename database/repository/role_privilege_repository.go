package repository

import (
	"errors"
	"fiber-poc-api/database/entity"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type RolePrivilegeRepository struct {
	gormDB *gorm.DB
}

func NewRolePrivilegeRepository(gormDB *gorm.DB) RolePrivilegeRepository {
	return RolePrivilegeRepository{gormDB}
}

func (repo RolePrivilegeRepository) FindByRoleCodeAndPrivilegeCode(roleCode, privilegeCode string) (*entity.RolePrivilege, error) {
	var privilege entity.RolePrivilege
	//privilege.RoleCode = roleCode
	//privilege.PrivilegeCode = privilegeCode
	err := repo.gormDB.Find(&privilege, "role_code = ? AND privilege_code = ?", roleCode, privilegeCode)
	if err.Error != nil {
		log.Errorf("FindByRoleCodeAndPrivilegeCode error: %+v", err.Error)
		return nil, errors.New(err.Error.Error())
	}
	if privilege.Id == 0 {
		return nil, nil
	}
	return &privilege, nil
}

func (repo RolePrivilegeRepository) CreateRolePrivilege(rolePrivilege *entity.RolePrivilege) error {
	err := repo.gormDB.Create(&rolePrivilege)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}
