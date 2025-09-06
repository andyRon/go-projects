package models

import "gorm.io/gorm"

// 某个用户进某个房间

type RoomUser struct {
	gorm.Model
	Rid uint `gorm:"column:rid;type:int(11);not null" json:"rid"`
	Uid uint `gorm:"column:uid;type:int(11);not null" json:"uid"`
}

func (table *RoomUser) TableName() string {
	return "room_user"
}
