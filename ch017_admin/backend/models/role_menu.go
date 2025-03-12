package models

import "gorm.io/gorm"

type RoleMenu struct {
	gorm.Model
	RoleId uint `gorm:"column:role_id;type:int(11);" json:"role_id"` // 角色ID
	MenuId uint `gorm:"column:menu_id;type:int(11);" json:"menu_id"` // 菜单ID
}

func (table *RoleMenu) TableName() string {
	return "role_menu"
}

// GetRoleMenus 获取角色菜单列表
func GetRoleMenus(roleIdentity string, isAdmin bool) (*gorm.DB, error) {
	tx := new(gorm.DB)
	if isAdmin {
		tx = DB.Model(new(Menu)).Select("id, ,identity, parent_id, path, name, sort, level, web_icon").Order("sort ASC")
	} else {
		role := new(Role)
		err := DB.Model(new(Role)).Select("id").Where("identity = ?", roleIdentity).Find(role).Error
		if err != nil {
			return nil, err
		}
		tx = DB.Model(new(RoleMenu)).
			Select("m.id, m,parent_id, m.identity, m.name, m.web_icon, m.sort, mb.path, m.level").
			Joins("left join menu m on role_menu.menu_id = menu.id").
			Where("role_menu.role_id =?", role.ID).
			Order("m.sort ASC")
	}
	return tx, nil
}

// GetRoleMenuIdentity 获取角色菜单标识
func GetRoleMenuIdentity(roleId uint, isAdmin bool) ([]string, error) {
	tx := new(gorm.DB)
	data := make([]string, 0)
	if isAdmin {
		tx = DB.Model(new(Menu)).Select("identity").Order("sort ASC")
	} else {
		tx = DB.Model(new(RoleMenu)).Select("m.identity").
			Joins("left join menu m on role_menu.menu_id = m.id").
			Where("role_menu.role_id =?", roleId).
			Order("m.sort ASC")
	}
	err := tx.Scan(&data).Error
	return data, err
}
