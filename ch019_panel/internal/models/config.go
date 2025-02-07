package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Key   string `gorm:"column:key; type:varchar(255)" json:"key"`
	Value string `gorm:"column:value; type:varchar(255)" json:"value"`
}

func (table *Config) TableName() string {
	return "config"
}
