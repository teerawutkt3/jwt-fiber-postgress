package repository

import (
	"errors"
	"fiber-poc-api/database/entity"
	"gorm.io/gorm"
)

type PrivilegeRepository struct {
	gormDB *gorm.DB
}

func NewPrivilegeRepository(gormDB *gorm.DB) PrivilegeRepository {
	return PrivilegeRepository{gormDB}
}

func (repo PrivilegeRepository) FindByCode(code string) (*entity.Privilege, error) {
	var privilege entity.Privilege
	privilege.Code = code
	err := repo.gormDB.First(&privilege)
	if err.Error != nil {
		return nil, errors.New(err.Error.Error())
	}
	return &privilege, nil
}

func (repo PrivilegeRepository) Save(role *entity.Privilege) error {
	err := repo.gormDB.Save(&role)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}
