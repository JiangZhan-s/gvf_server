package service

import (
	"fmt"
	"gorm.io/gorm"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/service/common"
	"gvf_server/utils"
	"path"
	"strconv"
	"strings"
	"time"
)

// CreateFile 添加文件数据
func CreateFile(filePath string, fileName string, fileSize int64, fId string, fileStoreId int, userID int) string {

	//获取文件后缀
	fileSuffix := path.Ext(fileName)
	//获取文件名
	filePrefix := fileName[0 : len(fileName)-len(fileSuffix)]
	fid, _ := strconv.Atoi(fId)
	//计算文件大小，并根据文件大小生成文件大小字符串
	var sizeStr string
	if fileSize < 1048576 {
		sizeStr = fmt.Sprintf("%dKB", fileSize/1024)
	} else {
		sizeStr = fmt.Sprintf("%dMB", fileSize/102400)
	}

	myFile := models.FileModel{
		FileName:       filePrefix,
		FileStoreID:    fileStoreId,
		FilePath:       filePath,
		UserID:         userID,
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderID: fid,
		Size:           fileSize / 1024,
		SizeStr:        sizeStr,
		Type:           utils.GetFileTypeInt(fileSuffix),
		Postfix:        strings.ToLower(fileSuffix),
		ShareFlag:      0,
	}
	// 将 myFile 保存到数据库中
	if err := global.DB.Create(&myFile).Error; err != nil {
		// 如果保存出错，输出错误信息并退出函数
		fmt.Println("failed to create file:", err)
		return ""
	}
	return fmt.Sprintf("%d", myFile.ID)

}

// GetUserFile 获取用户的文件
func GetUserFile(parentId string, storeId int) (files []models.FileModel) {
	global.DB.Find(&files, "file_store_id = ? and parent_folder_id = ?", storeId, parentId)
	return
}

// GetUserFileAll 获取用户的文件
func GetUserFileAll(storeId int, cr models.PageInfo) (files []models.FileModel, count int64, err error) {
	searchCond := "file_store_id = ?"
	searchValues := []interface{}{storeId}
	files, count, err = common.ComList(models.FileModel{}, common.Option{PageInfo: cr}, searchCond, searchValues...)
	return files, count, err
}

// SubtractSize 文件上传成功减去相应容量
func SubtractSize(size int64, fileStoreId int) {
	global.DB.Model(&models.FileStoreModel{}).Where("id = ?", fileStoreId).UpdateColumn("current_size", gorm.Expr("current_size + ?", size/1024))
}

// GetUserFileCount 获取用户文件数量
func GetUserFileCount(fileStoreId int) (fileCount int64) {
	var file []models.FileModel
	global.DB.Find(&file, "file_store_id = ?", fileStoreId).Count(&fileCount)
	return
}

// GetFileDetailUse 获取用户文件使用明细情况
func GetFileDetailUse(fileStoreId int) map[string]int64 {
	var files []models.FileModel
	var (
		docCount   int64
		imgCount   int64
		videoCount int64
		musicCount int64
		otherCount int64
	)

	fileDetailUseMap := make(map[string]int64, 0)

	//文档类型
	docCount = global.DB.Find(&files, "file_store_id = ? AND type = ?", fileStoreId, 1).RowsAffected
	fileDetailUseMap["docCount"] = docCount
	////图片类型
	imgCount = global.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 2).RowsAffected
	fileDetailUseMap["imgCount"] = imgCount
	//视频类型
	videoCount = global.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 3).RowsAffected
	fileDetailUseMap["videoCount"] = videoCount
	//音乐类型
	musicCount = global.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 4).RowsAffected
	fileDetailUseMap["musicCount"] = musicCount
	//其他类型
	otherCount = global.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 5).RowsAffected
	fileDetailUseMap["otherCount"] = otherCount

	return fileDetailUseMap
}

// GetTypeFile 根据文件类型获取文件
func GetTypeFile(fileType, fileStoreId int) (files []models.FileModel) {
	global.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, fileType)
	return
}

// CurrFileExists 判断当前文件夹是否有同名文件
func CurrFileExists(folderID, filename string) bool {
	var count int64
	//获取文件后缀
	fileSuffix := strings.ToLower(path.Ext(filename))
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	global.DB.Model(&models.FileModel{}).Where("parent_folder_id = ? AND file_name = ? AND postfix = ?", folderID, filePrefix, fileSuffix).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// GetFileInfo 通过fileId获取文件信息
func GetFileInfo(fId string) (file models.FileModel) {
	global.DB.First(&file, fId)
	return
}

// DownloadNumAdd 文件下载次数+1
func DownloadNumAdd(fileId string) {
	var file models.FileModel
	global.DB.First(&file, fileId)
	file.DownloadNum = file.DownloadNum + 1
	global.DB.Save(&file)
}

// DeleteUserFile 删除数据库文件数据
func DeleteUserFile(fId, folderId string, storeId int) {
	global.DB.Where("id = ? and file_store_id = ? and parent_folder_id = ?", fId, storeId, folderId).Delete(models.FileModel{})
}

// ShareFileUp 分享标志为1
func ShareFileUp(fileId string) {
	var file models.FileModel
	global.DB.First(&file, fileId)
	file.ShareFlag = 1
	global.DB.Save(&file)
}

// GetUserShareAll 获取用户的分享文件
func GetUserShareAll(storeId int, cr models.PageInfo) (files []models.FileModel, count int64, err error) {
	searchCond := "file_store_id = ? and share_flag=1"
	searchValues := []interface{}{storeId}
	files, count, err = common.ComList(models.FileModel{}, common.Option{PageInfo: cr}, searchCond, searchValues...)
	return files, count, err
}
