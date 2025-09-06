package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(100);uniqueIndex;not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(36);not null" json:"password"`
	Sdp      string `gorm:"column:sdp;type:text" json:"sdp"` // webrtc中重要概念，用于双方通信
}

func (table *User) TableName() string {
	return "user"
}
