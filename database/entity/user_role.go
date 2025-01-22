package entity

import "time"

type UserRole struct {
	Id          int64      `gorm:"primaryKey;column:id;AUTO_INCREMENT" json:"id"`
	UserId      int64      `gorm:"column:user_id;not null" json:"userId"`
	RoleCode    string     `gorm:"column:role_code;size:80;not null" json:"roleCode"`
	IsDeleted   string     `gorm:"column:is_deleted;size:1;not null" json:"isDeleted"`
	CreatedBy   *string    `gorm:"column:created_by;size:100" json:"createdBy"`
	UpdatedBy   *string    `gorm:"column:updated_by;size:100" json:"updatedBy"`
	CreatedDate time.Time  `gorm:"column:created_date;not null" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updated_date" json:"updatedDate"`
	Role        Role       `gorm:"references:Code;foreignKey:RoleCode"`
	User        User       `gorm:"references:Id;foreignKey:UserId"`
}

func (UserRole) TableName() string {
	return "user_role"
}
