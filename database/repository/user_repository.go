package repository

import (
	"errors"
	"fiber-poc-api/database/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) UserRepository {
	return UserRepository{gormDB}
}

func (repo UserRepository) GetUserAll() ([]entity.User, error) {
	users := new([]entity.User)
	gormDB := repo.gormDB.Find(users)
	if err := gormDB.Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return *users, nil
}

func (repo UserRepository) GetUserByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	gormDB := repo.gormDB.Where(entity.User{
		Username: username,
	}).First(user)
	if err := gormDB.Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, errors.New(err.Error())
	}
	return user, nil
}

func (repo UserRepository) CreateUser(user *entity.User) error {
	err := repo.gormDB.Create(&user)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}

func (repo UserRepository) UpdateUser(user *entity.User) error {
	//updateExpression := "SET #isDeleted = :isDeleted, #updatedDate = :updatedDate"
	//expressionAttributeNames := map[string]*string{
	//	"#isDeleted":   aws.String("isDeleted"),
	//	"#updatedDate": aws.String("updatedDate"),
	//}
	//expressionAttributeValues := map[string]*db.AttributeValue{
	//	":isDeleted": {
	//		S: aws.String("Y"),
	//	},
	//	":updatedDate": {
	//		S: aws.String(user.UpdatedDate),
	//	},
	//}
	//input := &db.UpdateItemInput{
	//	TableName: aws.String("user"),
	//	Key: map[string]*db.AttributeValue{
	//		"username": {
	//			S: aws.String("top"),
	//		},
	//	},
	//	UpdateExpression:          aws.String(updateExpression),
	//	ExpressionAttributeNames:  expressionAttributeNames,
	//	ExpressionAttributeValues: expressionAttributeValues,
	//	ReturnValues:              aws.String("UPDATED_NEW"), // Return the updated item
	//}
	//result, err := repo.db.UpdateItem(input)
	//if err != nil {
	//	log.Errorf("error PutItem error message: %+v", err.Error())
	//	return err
	//}
	//log.Infof("result: %+v", result)

	return nil
}

func (repo UserRepository) CreateUserRole(userRole *entity.UserRole) error {
	err := repo.gormDB.Create(&userRole)
	if err.Error != nil {
		return errors.New(err.Error.Error())
	}
	return nil
}

func (repo UserRepository) GetUserRole(userId int64, roleCode string) (*entity.UserRole, error) {
	userRole := &entity.UserRole{
		UserId:   userId,
		RoleCode: roleCode,
	}
	err := repo.gormDB.Find(&userRole)
	if err.Error != nil {
		return nil, errors.New(err.Error.Error())
	}
	if userRole.Id == 0 {
		return nil, nil
	}
	return userRole, nil
}
