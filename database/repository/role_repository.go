package repository

import (
	"errors"
	"fiber-poc-api/database/entity"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type RoleRepository struct {
	gormDB *gorm.DB
}

func NewRoleRepository(gormDB *gorm.DB) RoleRepository {
	return RoleRepository{gormDB}
}

func (repo RoleRepository) FindByCode(code string) (*entity.Role, error) {
	var role entity.Role
	role.Code = code
	// Perform the query
	err := repo.gormDB.First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("role with code '%s' not found", code)
		}
		return nil, fmt.Errorf("failed to find role by code: %w", err)
	}
	return &role, nil
}

func (repo RoleRepository) Create(role *entity.Role) error {
	err := repo.gormDB.Create(&role)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}

func (repo RoleRepository) Update(role *entity.Role) error {
	err := repo.gormDB.Save(role)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}

func (repo RoleRepository) GetRolePrivilege(roleCode, privilegeCode string) (*entity.RolePrivilege, error) {
	rolePrivilege := &entity.RolePrivilege{
		RoleCode:      roleCode,
		PrivilegeCode: privilegeCode,
	}
	err := repo.gormDB.Where("role_code = ? AND privilege_code = ?", roleCode, privilegeCode).First(&rolePrivilege)
	if err.Error != nil {
		log.Errorf("GetRolePrivilege error: %+v", err.Error)
		return nil, errors.New(err.Error.Error())
	}
	if rolePrivilege.Id == 0 {
		return nil, nil
	}
	return rolePrivilege, nil
}

func (repo RoleRepository) CreateRolePrivilege(rolePrivilege *entity.RolePrivilege) error {
	err := repo.gormDB.Create(&rolePrivilege)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}
