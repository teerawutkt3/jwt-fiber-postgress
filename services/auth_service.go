package services

import (
	"errors"
	"fiber-poc-api/constant"
	"fiber-poc-api/database/entity"
	"fiber-poc-api/database/repository"
	"fiber-poc-api/model"
	"fiber-poc-api/utils"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type AuthService struct {
	userRepo         repository.UserRepository
	loginHistoryRepo repository.LoginHistoryRepository
}

func NewAuthService(userRepo repository.UserRepository, loginHistoryRepo repository.LoginHistoryRepository) AuthService {
	return AuthService{userRepo, loginHistoryRepo}
}

func (s AuthService) Login(req model.LoginReq, xRequestId string) (*string, error) {

	user, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		log.Infof("[%s] error during GetUserByUsername err: %+v", xRequestId, err.Error())
		errDesc := "error GetUserByUsername"
		s.createLoginTransaction(req.Username, constant.LOGIN_STATUS_FAIL, &errDesc)
		return nil, err
	}

	err = utils.ValidatePassword(user.Password, req.Password)
	if err != nil {
		log.Infof("[%s] ValidatePassword not pass, UNAUTHORIZED", xRequestId)
		errDesc := "ValidatePassword"
		s.createLoginTransaction(req.Username, constant.LOGIN_STATUS_FAIL, &errDesc)
		return nil, errors.New("UNAUTHORIZED")
	}

	expTime := viper.GetDuration("jwt.expire")
	claims := jwt.MapClaims{
		"username": user.Username,
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * expTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		errDesc := err.Error()
		s.createLoginTransaction(req.Username, constant.LOGIN_STATUS_FAIL, &errDesc)
		return nil, err
	}

	s.createLoginTransaction(req.Username, constant.LOGIN_STATUS_SUCCESS, nil)
	return &t, nil
}

func (s AuthService) createLoginTransaction(username, status string, errDesc *string) {
	var history = entity.LoginHistory{}
	history.Username = username
	history.Status = status
	history.IsDeleted = "N"
	history.ErrorDesc = errDesc
	history.CreatedDate = time.Now()
	err := s.loginHistoryRepo.Create(history)
	if err != nil {
		log.Errorf("error during craete LoginHistory err: %+v", err.Error())
	}
}

func (s AuthService) Register(req model.LoginReq, xRequestId string) error {

	// ==> validate user exits
	userExit, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		log.Infof("[%s] error during GetUserByUsername err: %+v", xRequestId, err.Error())
		return err
	}
	if userExit != nil {
		msg := "username exits"
		log.Errorf("[%s] error register %s %s", xRequestId, req.Username, msg)
		return errors.New(msg)
	}

	// ==> create user
	dateNow := time.Now()
	password, _ := utils.HashPassword(req.Password)
	user := &entity.User{
		Username:    req.Username,
		Password:    password,
		IsDeleted:   "N",
		CreatedDate: dateNow,
		UpdatedDate: dateNow,
	}
	err = s.userRepo.CreateUser(user)
	if err != nil {
		log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		return err
	}
	return nil
}

func (s AuthService) UpdateUser(req model.LoginReq, xRequestId string) error {

	user := &entity.User{
		Username:    req.Username,
		Password:    req.Password,
		IsDeleted:   "N",
		UpdatedDate: time.Now(),
	}
	err := s.userRepo.UpdateUser(user)
	if err != nil {
		log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		return err
	}
	return nil
}

func (s AuthService) GetUserAll(xRequestId string) []entity.User {

	log.Infof("[%s] GetUserAll", xRequestId)
	users, err := s.userRepo.GetUserAll()
	if err != nil {
		log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		return make([]entity.User, 0)
	}
	return users
}
