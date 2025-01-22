package entity

import "time"

type RolePrivilege struct {
	Id            int64     `gorm:"primaryKey;column:id;AUTO_INCREMENT" json:"id"`
	PrivilegeCode string    `gorm:"column:privilege_code;not null" json:"privilegeCode"`
	RoleCode      string    `gorm:"column:role_code;size:80;not null" json:"roleCode"`
	IsDeleted     string    `gorm:"column:is_deleted;size:1;not null" json:"isDeleted"`
	CreatedDate   time.Time `gorm:"column:created_date;not null" json:"createdDate"`
	UpdatedDate   time.Time `gorm:"column:updated_date;not null" json:"updatedDate"`
	Role          Role      `gorm:"references:Code;foreignKey:RoleCode"`
	Privilege     Privilege `gorm:"references:Code;foreignKey:PrivilegeCode"`
}

func (RolePrivilege) TableName() string {
	return "role_privilege"
}
