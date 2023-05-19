package service

import (
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/service/common"
)

// GetRoleAll 获取所有角色
func GetRoleAll(cr models.PageInfo) (roles []models.RoleModel, count int64, err error) {
	searchCond := ""
	var searchValues []interface{}
	roles, count, err = common.ComList(models.RoleModel{}, common.Option{PageInfo: cr}, searchCond, searchValues...)
	return roles, count, err
}

// AddRole 添加角色
func AddRole(role models.RoleModel) (err error) {

	//入库
	err = global.DB.Create(&role).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}
	return err
}

func DeleteRoleById(roleID uint) (string, error) {
	// 查询要删除的角色
	var role models.RoleModel
	err := global.DB.First(&role, roleID).Error
	if err != nil {
		global.Log.Error(err)
		return "", err
	}

	// 删除角色
	err = global.DB.Delete(&role).Error
	if err != nil {
		global.Log.Error(err)
		return "", err
	}

	return role.RoleName, nil
}

func GetPermissionAll() (permissions []models.PermissionModel, err error) {
	err = global.DB.Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, err
}

func DeleteRolePermissionByRoleId(roleId uint) error {
	// 先删除该角色的所有权限映射
	if err := global.DB.Where("role_id = ?", roleId).Delete(&models.RolePermissionModel{}).Error; err != nil {
		return err
	}
	return nil
}

func AddRolePermissionBatch(roleId uint, permissionIds []uint) error {
	var rolePermissions []models.RolePermissionModel
	for _, permissionID := range permissionIds {
		rolePermission := models.RolePermissionModel{
			RoleID:       roleId,
			PermissionID: permissionID,
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	if err := global.DB.Create(&rolePermissions).Error; err != nil {
		return err
	}
	return nil
}
