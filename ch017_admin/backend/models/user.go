package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity     string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	RoleIdentity string `gorm:"column:role_identity;type:varchar(36);" json:"role_identity"` // 角色唯一标识
	Role         Role   `gorm:"foreignKey:identity;references:role_identity;" json:"role"`   // 管理角色基础表
	Username     string `gorm:"column:username;type:varchar(50);" json:"username"`
	Password     string `gorm:"column:password;type:varchar(36);" json:"password"`
	Phone        string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId    string `gorm:"column:wx_union_id;type:varchar(255);" json:"wx_union_id"`
	WxOpenId     string `gorm:"column:wx_open_id;type:varchar(255);" json:"wx_open_id"`
	Avatar       string `gorm:"column:avatar;type:varchar(255);" json:"avatar"` // 头像
}

func (table *User) TableName() string {
	return "user"
}

// TODO
