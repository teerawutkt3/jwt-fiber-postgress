package handlers

import (
	"fiber-poc-api/common"
	"fiber-poc-api/model"
	"fiber-poc-api/services"
	"fiber-poc-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type AuthHandler struct {
	svc services.AuthService
}

func NewAuthHandler(svc services.AuthService) AuthHandler {
	return AuthHandler{svc}
}

func (h AuthHandler) LoginHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := model.LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Login invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] username %s login date: %+v", xRequestId, req.Username, time.Now())

	token, err := h.svc.Login(req, xRequestId)
	if err != nil {
		log.Errorf("[%s] Error during login err: %+v", xRequestId, err.Error())
		if "UNAUTHORIZED" == err.Error() {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	log.Infof("[%s] login success", xRequestId)
	response := model.LoginRes{}
	response.Token = *token
	return c.JSON(response)
}

func (h AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := model.LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Register invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] Req username: %+v", xRequestId, req.Username)
	err = h.svc.Register(req, xRequestId)
	if err != nil {
		log.Errorf("[%s] Register Error", xRequestId)
		return c.Status(500).JSON(fiber.Map{"message": "error"})
	}
	log.Infof("[%s] Register success", xRequestId)
	response := common.Response{
		Message: "success",
	}
	return c.JSON(response)
}

func (h AuthHandler) GetUserAllHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := model.LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] GetUserHandler invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	log.Infof("[%s] Req: %+v", xRequestId, req)
	users := h.svc.GetUserAll(xRequestId)
	log.Infof("[%s] Register success", xRequestId)
	return c.JSON(users)
}
