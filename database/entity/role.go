package entity

import "time"

type Role struct {
	Code        string    `gorm:"primaryKey;column:code;unique;size:80" json:"code"`
	Name        string    `gorm:"unique;column:name;unique;size:100:not null" json:"name"`
	IsDeleted   string    `gorm:"column:is_deleted;size:1;not null" json:"isDeleted"`
	CreatedDate time.Time `gorm:"column:created_date;not null" json:"createdDate"`
	UpdatedDate time.Time `gorm:"column:updated_date;not null" json:"updatedDate"`
}

func (Role) TableName() string {
	return "role"
}
