package models

import "gorm.io/gorm"

// RoleModel 角色表
type RoleModel struct {
	gorm.Model
	RoleName    string `json:"role_name"` //文件夹名
	Description string `json:"description"`
}

// PermissionModel 权限表
type PermissionModel struct {
	gorm.Model
	PermissionName string `json:"permission_name"` //文件夹名
	Description    string `json:"description"`
}

// RolePermissionModel 角色-权限映射表
type RolePermissionModel struct {
	gorm.Model
	RoleID       uint            `json:"role_id" gorm:"foreignKey:RoleID"`               // 外键，关联RoleModel的ID字段
	PermissionID uint            `json:"permission_id" gorm:"foreignKey:PermissionID"`   // 外键，关联PermissionModel的ID字段
	Role         RoleModel       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 角色关联
	Permission   PermissionModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 权限关联
}

// UserRoleModel 用户-角色映射表
type UserRoleModel struct {
	gorm.Model
	RoleID uint      `json:"role_id" gorm:"foreignKey:RoleID"`               // 外键，关联RoleModel的ID字段
	UserID uint      `json:"user_id" gorm:"foreignKey:UserID"`               // 外键，关联UserModel的ID字段
	Role   RoleModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 角色关联
	User   UserModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 用户关联
}
