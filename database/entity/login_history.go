package entity

import "time"

type LoginHistory struct {
	Id          int64     `gorm:"primaryKey;column:id;AUTO_INCREMENT" json:"id"`
	Username    string    `gorm:"column:username;index;size:200;not null" json:"username"`
	Status      string    `gorm:"column:status;size:7;not null" json:"status"` // SUCCESS, FAIL
	ErrorDesc   *string   `gorm:"column:error_desc;size:100" json:"ERROR_DESC"`
	IsDeleted   string    `gorm:"column:is_deleted;size:1;not null" json:"isDeleted"`
	CreatedDate time.Time `gorm:"column:created_date;not null" json:"createdDate"`
}

func (LoginHistory) TableName() string {
	return "login_history"
}
