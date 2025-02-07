package models

import "gorm.io/gorm"

type Web struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"` // 唯一标识
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`
	Domain   string `gorm:"column:domain;type:varchar(255);" json:"domain"`
	Dir      string `gorm:"column:dir;type:varchar(255);" json:"dir"`             // 静态文件目录
	ConfPath string `gorm:"column:conf_path;type:varchar(255);" json:"conf_path"` // 配置文件路径
}

func (table Web) TableName() string {
	return "web"
}
