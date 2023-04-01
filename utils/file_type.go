package utils

import "strings"

// GetFileTypeInt 判断文件后缀获取类型id
func GetFileTypeInt(filePrefix string) int {
	filePrefix = strings.ToLower(filePrefix)
	if filePrefix == ".doc" || filePrefix == ".docx" || filePrefix == ".txt" || filePrefix == ".pdf" {
		return 1
	}
	if filePrefix == ".jpg" || filePrefix == ".png" || filePrefix == ".gif" || filePrefix == ".jpeg" {
		return 2
	}
	if filePrefix == ".mp4" || filePrefix == ".avi" || filePrefix == ".mov" || filePrefix == ".rmvb" || filePrefix == ".rm" {
		return 3
	}
	if filePrefix == ".mp3" || filePrefix == ".cda" || filePrefix == ".wav" || filePrefix == ".wma" || filePrefix == ".ogg" {
		return 4
	}

	return 5
}
