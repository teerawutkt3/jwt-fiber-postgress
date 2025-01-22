package repository

import (
	"errors"
	"fiber-poc-api/database/entity"
	"gorm.io/gorm"
)

type LoginHistoryRepository struct {
	gormDB *gorm.DB
}

func NewLoginHistoryRepository(gormDB *gorm.DB) LoginHistoryRepository {

	return LoginHistoryRepository{gormDB}
}

func (repo LoginHistoryRepository) ListAll() ([]entity.LoginHistory, error) {
	historyList := new([]entity.LoginHistory)
	gormDB := repo.gormDB.Find(historyList)
	if err := gormDB.Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return *historyList, nil
}

func (repo LoginHistoryRepository) Create(history entity.LoginHistory) error {
	err := repo.gormDB.Create(&history)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}
