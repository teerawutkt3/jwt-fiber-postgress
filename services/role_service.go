package services

import (
	"fiber-poc-api/constant"
	"fiber-poc-api/database/entity"
	"fiber-poc-api/database/repository"
	"fiber-poc-api/model"
	"fiber-poc-api/utils"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type RoleService struct {
	userRepo          repository.UserRepository
	roleRepo          repository.RoleRepository
	privilegeRepo     repository.PrivilegeRepository
	rolePrivilegeRepo repository.RolePrivilegeRepository
}

func NewRoleService(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
	privilegeRepo repository.PrivilegeRepository,
	rolePrivilegeRepo repository.RolePrivilegeRepository,
) RoleService {
	return RoleService{
		userRepo,
		roleRepo,
		privilegeRepo,
		rolePrivilegeRepo,
	}
}

func (service RoleService) CreateRole() error {
	return nil
}

func (service RoleService) CreateRolePrivileges(req model.CreateRolePrivilegesReq) error {
	log.Infof("CreateRolePrivileges req: %+v", req)
	roleCode := req.RoleCode
	privileges := req.Privileges
	for _, privilege := range privileges {
		rolePrivilege, err := service.rolePrivilegeRepo.FindByRoleCodeAndPrivilegeCode(roleCode, privilege)
		if err != nil {
			return err
		}
		if rolePrivilege == nil {
			err := service.rolePrivilegeRepo.CreateRolePrivilege(&entity.RolePrivilege{
				RoleCode:      roleCode,
				PrivilegeCode: privilege,
			})
			if err != nil {
				return err
			}
		}
	}
	log.Infof("CreateRolePrivileges completed req: %+v", req)
	return nil
}

func (service RoleService) Initial(xRequestId string) {

	log.Infof("[%s] Initial", xRequestId)

	now := time.Now()

	// create role
	var roles = constant.Roles()
	for i := 0; i < len(roles); i++ {
		code := roles[i][0]
		name := roles[i][1]
		roleEntity, _ := service.roleRepo.FindByCode(code)
		if roleEntity == nil {
			err := service.roleRepo.Create(&entity.Role{
				Code:        code,
				Name:        name,
				IsDeleted:   "N",
				CreatedDate: now,
				UpdatedDate: now,
			})

			if err != nil {
				log.Errorf("[%s] CreateRole code: %s error: %+v", xRequestId, code, err.Error())
			}
		} else {
			err := service.roleRepo.Update(&entity.Role{
				Code:        code,
				Name:        name,
				IsDeleted:   "N",
				CreatedDate: now,
				UpdatedDate: now,
			})

			if err != nil {
				log.Errorf("[%s] UpdateRole code: %s error: %+v", xRequestId, code, err.Error())
			}
		}
	}

	// create privilege
	var privileges = constant.Privileges()
	for i := 0; i < len(privileges); i++ {
		privilegeGroupCode := privileges[i][0]
		privilegeGroupName := privileges[i][1]
		privilegeCode := privileges[i][2]
		privilegeName := privileges[i][3]

		privilegeEntity, _ := service.privilegeRepo.FindByCode(privilegeCode)
		if privilegeEntity == nil {
			err := service.privilegeRepo.Save(&entity.Privilege{
				Code:        privilegeCode,
				Name:        privilegeName,
				GroupCode:   privilegeGroupCode,
				GroupName:   privilegeGroupName,
				IsDeleted:   "N",
				CreatedDate: now,
				UpdatedDate: &now,
			})

			if err != nil {
				log.Errorf("[%s] CreatePrivilege code: %s error: %+v", xRequestId, privilegeCode, err.Error())
			}
		}
	}

	// mapping role SUPER_ADMIN
	for i := 0; i < len(privileges); i++ {
		privilegeCode := privileges[i][2]
		rolePrivilege, _ := service.roleRepo.GetRolePrivilege(constant.SUPER_ADMIN, privilegeCode)
		if rolePrivilege == nil {
			err := service.roleRepo.CreateRolePrivilege(&entity.RolePrivilege{
				RoleCode:      constant.SUPER_ADMIN,
				PrivilegeCode: privilegeCode,
				IsDeleted:     "N",
				CreatedDate:   now,
				UpdatedDate:   now,
			})
			if err != nil {
				log.Errorf("[%s] SaveRolePrivilege error: %+v", xRequestId, err.Error())
			}
		}
	}

	// create user
	username := "superadmin"
	password := "superadmin"
	passwordEncrypt, _ := utils.HashPassword(password)

	user, _ := service.userRepo.GetUserByUsername(username)
	if user == nil {
		err := service.userRepo.CreateUser(&entity.User{
			Username:    username,
			Password:    passwordEncrypt,
			IsDeleted:   "N",
			CreatedDate: now,
		})
		if err != nil {
			log.Errorf("[%s] CreateUser error: %+v", xRequestId, err.Error())
		}
	}

	// Mapping user role
	userRole, _ := service.userRepo.GetUserRole(user.Id, constant.SUPER_ADMIN)
	if userRole == nil {
		err := service.userRepo.CreateUserRole(&entity.UserRole{
			UserId:      user.Id,
			RoleCode:    constant.SUPER_ADMIN,
			IsDeleted:   "N",
			CreatedDate: now,
		})
		if err != nil {
			log.Errorf("[%s] SaveUserRole error: %+v", xRequestId, err.Error())
		}
	}
	log.Infof("[%s] Initial Complete", xRequestId)
}
