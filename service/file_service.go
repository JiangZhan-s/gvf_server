package service

import (
	"fmt"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/utils"
	"path"
	"strconv"
	"strings"
	"time"
)

// CreateFile 添加文件数据
func CreateFile(fileName string, fileSize int64, fId string, fileStoreId int, userID int) string {
	var sizeStr string
	//获取文件后缀
	fileSuffix := path.Ext(fileName)
	//获取文件名
	filePrefix := fileName[0 : len(fileName)-len(fileSuffix)]
	fid, _ := strconv.Atoi(fId)

	if fileSize < 1048576 {
		sizeStr = strconv.FormatInt(fileSize/1024, 10) + "KB"
	} else {
		sizeStr = strconv.FormatInt(fileSize/102400, 10) + "MB"
	}

	myFile := models.FileModel{
		FileName:       filePrefix,
		FileStoreID:    fileStoreId,
		FilePath:       "",
		UserID:         userID,
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderID: fid,
		Size:           fileSize / 1024,
		SizeStr:        sizeStr,
		Type:           utils.GetFileTypeInt(fileSuffix),
		Postfix:        strings.ToLower(fileSuffix),
	}
	// 将 myFile 保存到数据库中
	if err := global.DB.Create(&myFile).Error; err != nil {
		// 如果保存出错，输出错误信息并退出函数
		fmt.Println("failed to create file:", err)
		return ""
	}
	return string(myFile.ID)

}

// GetUserFile 获取用户的文件
func GetUserFile(parentId string, storeId int) (files []models.FileModel) {
	global.DB.Find(&files, "file_store_id = ? and parent_folder_id = ?", storeId, parentId)
	return
}

// SubtractSize 文件上传成功减去相应容量
func SubtractSize(size int64, fileStoreId int) {
	var fileStore models.FileStoreModel
	fileStore.CurrentSize = fileStore.CurrentSize + size/1024
	fmt.Println(fileStore.CurrentSize)
	global.DB.Save(&fileStore)
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
func CurrFileExists(fId, filename string) bool {
	var file models.FileModel
	//获取文件后缀
	fileSuffix := strings.ToLower(path.Ext(filename))
	//获取文件名
	filePrefix := filename[0 : len(filename)-len(fileSuffix)]

	global.DB.Find(&file, "parent_folder_id = ? and file_name = ? and postfix = ?", fId, filePrefix, fileSuffix)

	if file.Size > 0 {
		return false
	}
	return true
}

// GetFileInfo 通过fileId获取文件信息
func GetFileInfo(fId string) (file models.FileModel) {
	global.DB.First(&file, fId)
	return
}

// DownloadNumAdd 文件下载次数+1
func DownloadNumAdd(fId string) {
	var file models.FileModel
	global.DB.First(&file, fId)
	file.DownloadNum = file.DownloadNum + 1
	global.DB.Save(&file)
}

// DeleteUserFile 删除数据库文件数据
func DeleteUserFile(fId, folderId string, storeId int) {
	global.DB.Where("id = ? and file_store_id = ? and parent_folder_id = ?", fId, storeId, folderId).Delete(models.FileModel{})
}
