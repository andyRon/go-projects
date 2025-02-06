package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(100);uniqueIndex;not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(36);not null" json:"password"`
	Sdp      string `gorm:"column:sdp;type:text" json:"sdp"`
}

func (table *User) TableName() string {
	return "user"
}
