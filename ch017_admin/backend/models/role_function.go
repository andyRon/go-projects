package models

import (
	"github.com/andyron/mini-admin/helper"
	"gorm.io/gorm"
)

type RoleFunction struct {
	gorm.Model
	RoleId     uint `gorm:"column:role_id;type:int(11);" json:"role_id"`         // 角色ID
	FunctionId uint `gorm:"column:function_id;type:int(11);" json:"function_id"` // 功能ID
}

func (table *RoleFunction) TableName() string {
	return "role_function"
}

type authFuncReply struct {
	URI string `json:"uri"`
}

// GetAuthFunctionURI 获取角色授权功能的URI
func GetAuthFunctionURI(roleIdentity string) (map[string]interface{}, error) {
	role := new(Role)
	err := DB.Model(new(Role)).Select("id").Where("identity =?", roleIdentity).Find(role).Error
	if err != nil {
		return nil, err
	}
	afr := make([]*authFuncReply, 0)
	err = DB.Model(new(RoleFunction)).Select("f.uri").
		Joins("Left Join function f ON f.id = role_function.function_id").
		Where("role_function.role_id = ?", role.ID).Find(&afr).Error
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	for _, v := range afr {
		data[v.URI] = "1"
	}
	return data, err
}

// GetRoleFunctionIdentity 获取角色功能标识
func GetRoleFunctionIdentity(roleId uint, isAdmin bool) ([]string, error) {
	tx := new(gorm.DB)
	data := make([]string, 0)
	if isAdmin {
		tx = DB.Model(new(Function)).Select("identity").Order("sort ASC")
	} else {
		tx = DB.Model(new(RoleFunction)).Select("f.identity").
			Joins("Left Join function f ON f.id = role_function.function_id").
			Where("role_function.role_id =?", roleId).Order("f.sort ASC")
	}
	err := tx.Scan(&data).Error
	return data, err
}

// GetRoleFunctions 获取角色功能列表
func GetRoleFunctions(roleIdentity string, isAdmin bool) *gorm.DB {
	tx := DB.Model(new(Function)).Select("function.identity, m.identity menu_identity, function.name, function.uri, function.sort").
		Joins("Left Join menu m ON m.id = function.menu_id")
	if !isAdmin {
		var roleId int
		err := DB.Model(new(Role)).Select("id").Where("identity =?", roleIdentity).Scan(&roleId).Error
		if err != nil {
			helper.Error("[DB ERROR] get role function err: %v", err)
		}
		tx.Joins("Left Join role_function rf ON rf.function_id = function.id").Where("rf.role_id =?", roleId)
	}
	tx.Order("function.sort ASC")
	return tx
}
