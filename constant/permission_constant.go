package constant

// Roles Code
const SUPER_ADMIN = "SUPER_ADMIN"
const COMPANY = "COMPANY"
const SEALER = "SEALER"

// Privileges Code
const CAN_ACCESS_MENU_USER = "CAN_ACCESS_MENU_USER"
const CAN_CREATE_USER = "CAN_CREATE_USER"
const CAN_EDIT_USER = "CAN_EDIT_USER"
const CAN_GET_USER = "CAN_GET_USER"
const CAN_DELETE_USER = "CAN_DELETE_USER"
const CAN_READ_ROLE = "CAN_READ_ROLE"
const CAN_CREATE_ROLE = "CAN_CREATE_ROLE"
const CAN_DELETE_ROLE = "CAN_DELETE_ROLE"
const CAN_GET_PRIVILEGE = "CAN_GET_PRIVILEGE"
const CAN_UPLOAD = "CAN_UPLOAD"

func Roles() [][]string {
	return [][]string{
		{SUPER_ADMIN, "Super Admin"},
		{COMPANY, "Company"},
		{SEALER, "Sealer"},
	}
}

func Privileges() [][]string {
	// GROUP_CODE | GROUP_NAME | PRIVILEGE_CODE | PRIVILEGE_NAME
	return [][]string{
		{"USER", "User", CAN_ACCESS_MENU_USER, "Access menu user"},
		{"USER", "User", CAN_CREATE_USER, "Create user"},
		{"USER", "User", CAN_EDIT_USER, "Edit user"},
		{"USER", "User", CAN_GET_USER, "Get user"},
		{"USER", "User", CAN_DELETE_USER, "Delete user"},

		{"ROLE", "Role", CAN_READ_ROLE, "Read role"},
		{"ROLE", "Role", CAN_CREATE_ROLE, "Create role"},
		{"ROLE", "Role", CAN_DELETE_ROLE, "Delete role"},
		{"ROLE", "Role", CAN_GET_PRIVILEGE, "Get privilege"},

		{"UPLOAD", "Upload", CAN_UPLOAD, "upload file"},
	}
}
