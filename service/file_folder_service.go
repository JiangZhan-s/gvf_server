package service

import (
	"errors"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/service/common"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// CreateFolderRoot 创建用户根目录
// 根目录名称为用户的用户名，父目录ID为0
func CreateFolderRoot(fileStoreId int, userName string) (*models.FileFolderModel, error) {
	fileFolder := models.FileFolderModel{
		FileFolderName: userName,
		FileStoreID:    fileStoreId,
		ParentFolderID: 0,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := global.DB.Create(&fileFolder).Error; err != nil {
		return nil, err
	}
	folderPath := filepath.Join(global.Path, userName)
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		panic(err)
	}
	return &fileFolder, nil
}

// FindFolderRoot 查询用户根目录
// 根据用户所对应的仓库id查询用户的根目录(根目录的父id为0)
func FindFolderRoot(fileStoreId int) (int, error) {
	result := -1
	if err := global.DB.Model(&models.FileFolderModel{}).Select("id").Where("file_store_id=? and parent_folder_id=?", fileStoreId, 0).Scan(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

// CreateFolder 新建文件夹
func CreateFolder(folderName string, parentId int, fileStoreId int) (models.FileFolderModel, error) {
	fileFolder := models.FileFolderModel{
		FileFolderName: folderName,
		ParentFolderID: parentId,
		FileStoreID:    fileStoreId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	global.DB.Create(&fileFolder)
	// 创建文件夹
	folderPath := GetCurrentFolderPath(fileFolder)
	err := os.MkdirAll(global.Path+"/"+folderPath, 0755)
	if err != nil {
		return models.FileFolderModel{}, err
	}
	return fileFolder, err
}

// GetCurrentFolderPath 获取当前路径所有的父级
func GetCurrentFolderPath(folder models.FileFolderModel) string {
	var parentFolder models.FileFolderModel
	if folder.ParentFolderID != 0 {
		global.DB.Find(&parentFolder, "id = ?", folder.ParentFolderID)
		path := "/" + folder.FileFolderName
		//递归查找当前所有父级
		return GetCurrentFolderPath(parentFolder) + path
	}
	//最顶层文件夹
	return folder.FileFolderName
}

// GetFolderByParent 获取当前目录下的文件夹
func GetFolderByParent(storeId int, cr models.PageInfo) (files []models.FileModel, count int64, err error) {
	searchCond := "file_store_id = ?"
	searchValues := []interface{}{storeId}
	files, count, err = common.ComList(models.FileModel{}, common.Option{PageInfo: cr}, searchCond, searchValues...)
	return files, count, err
}

// GetFolderById 根据ID获取文件夹
func GetFolderById(folderId string) (folder models.FileFolderModel, err error) {
	err = global.DB.Where("id = ? ", folderId).First(&folder).Error
	if err != nil {
		return folder, err
	}

	return folder, nil
}

// GetParentFolderIDById 根据ID获取父文件夹ID
func GetParentFolderIDById(folderId string) (parentId int, err error) {
	var folder models.FileFolderModel
	err = global.DB.Where("id = ?", folderId).First(&folder).Error
	if err != nil {
		self, _ := strconv.Atoi(folderId)
		return self, err
	}

	return folder.ParentFolderID, nil
}

// CurrFolderExists 判断当前目录下是否有同名文件夹
func CurrFolderExists(parentFolderID, folderName string) bool {
	var count int64

	global.DB.Model(&models.FileFolderModel{}).Where("parent_folder_id = ? AND file_folder_name = ?", parentFolderID, folderName).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func DeleteFolderById(folderId int) error {

	// 检查是否存在具有给定文件夹ID作为父文件夹ID的子文件夹
	var count int64
	global.DB.Model(&models.FileFolderModel{}).Where("parent_folder_id = ?", folderId).Count(&count)
	if count > 0 {
		return errors.New("该文件夹包含子文件夹，无法删除")
	}
	global.DB.Model(&models.FileModel{}).Where("parent_folder_id = ?", folderId).Count(&count)
	if count > 0 {
		return errors.New("该文件夹包含子文件，无法删除")
	}
	var folder models.FileFolderModel
	result := global.DB.First(&folder, folderId)
	if result.Error != nil {
		return result.Error
	}

	// 执行删除操作
	if err := global.DB.Delete(&folder).Error; err != nil {
		return err
	}

	return nil
}

func GetFolderCount(storeId int) (count int64, err error) {
	err = global.DB.Model(&models.FileFolderModel{}).Where("file_store_id = ?", storeId).Count(&count).Error
	return
}
