package handlers

import (
	"fiber-poc-api/model"
	"fiber-poc-api/services"
	"fiber-poc-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type RoleHandler struct {
	svc services.RoleService
}

func NewRoleHandler(svc services.RoleService) RoleHandler {
	return RoleHandler{svc}
}

func (h RoleHandler) CreateRoleHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := model.LoginReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] Login invalid Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	err = h.svc.CreateRole()
	if err != nil {
		return err
	}
	response := model.LoginRes{}
	//response.Token = *token
	return c.JSON(response)
}

func (h RoleHandler) CreateRolePrivilegeHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	req := model.CreateRolePrivilegesReq{}
	err := c.BodyParser(&req)
	if err != nil {
		log.Errorf("[%s] CreateRolePrivilegeHandler Request err: %+v", xRequestId, err.Error())
		return c.Status(400).JSON("invalid request")
	}
	err = h.svc.CreateRolePrivileges(req)
	if err != nil {
		return c.Status(500).JSON("error create role privileges")
	}

	return c.JSON("success")
}

func (h RoleHandler) InitialHandler(c *fiber.Ctx) error {
	xRequestId := utils.GetXRequestId()
	h.svc.Initial(xRequestId)
	return c.JSON("initial complete")
}
