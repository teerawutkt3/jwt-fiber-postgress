package entity

import "time"

type User struct {
	Id          int64     `gorm:"primaryKey;column:id;AUTO_INCREMENT" json:"id"`
	Username    string    `gorm:"column:username;index;unique;size:200;not null" json:"username"`
	Password    string    `gorm:"column:password;size:200;not null" json:"password"`
	IsDeleted   string    `gorm:"column:is_deleted;size:1;not null" json:"isDeleted"`
	CreatedDate time.Time `gorm:"column:created_date;not null" json:"createdDate"`
	UpdatedDate time.Time `gorm:"column:updated_date;not null" json:"updatedDate"`
}

func (User) TableName() string {
	return "user"
}
