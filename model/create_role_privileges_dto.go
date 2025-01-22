package model

type CreateRolePrivilegesReq struct {
	RoleCode   string   `json:"roleCode"`
	Privileges []string `json:"privileges"`
}
