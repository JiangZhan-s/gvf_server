package service

import (
	"fmt"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/utils"
	"strconv"
)

func CreateShare(fileId string, userId string) string {
	hash := utils.GetSHA256ByteHashCode([]byte(fileId + "file+user" + userId))
	fid, _ := strconv.Atoi(fileId)
	uid, _ := strconv.Atoi(userId)
	if GetShareByHash(hash).ID != 0 {
		return hash
	}

	myShare := models.ShareModel{
		FileID: fid,
		UserID: uid,
		Hash:   hash,
	}
	if err := global.DB.Create(&myShare).Error; err != nil {
		// 如果保存出错，输出错误信息并退出函数
		fmt.Println("failed to create share:", err)
	}
	return hash
}

func GetShareByHash(hash string) models.ShareModel {
	var share models.ShareModel
	global.DB.First(&share, "hash = ?", hash)
	return share
}
