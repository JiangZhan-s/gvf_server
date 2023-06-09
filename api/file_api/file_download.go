package file_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils"
	"gvf_server/utils/jwts"
	"gvf_server/utils/valid"
	"os"
	"path"
	"time"
)

func (FileApi) FileDownloadByNewView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fmt.Println(user.IP)

	fileId := c.GetHeader("file_id")
	fileModel := models.FileModel{}
	fileModel = service.GetFileInfo(fileId)
	filePath := path.Join(global.Path, fileModel.FilePath, fileModel.FileName+fileModel.Postfix)
	fileName := fileModel.FileName + fileModel.Postfix
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		c.Status(500)
		return
	}
	defer file.Close()
	fmt.Println("文件打开成功")

	hashData := utils.GetSHA256HashCode(file)
	fmt.Println(hashData)
	//对比链上连下数据哈希是否相等
	maxRetry := 5
	var msg string
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.QueryDataHash(fileId)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
	}

	//解析json数据
	var d models.Data
	err = json.Unmarshal([]byte(msg), &d)
	if err != nil {
		global.Log.Error(err)
		c.Status(500)
		panic(err)
	}

	if d.DataHash != hashData {
		err = errors.New("链上连下数据协同验证不统一，下载请求被拒绝")
		c.Status(550)
		return
	}
	fmt.Println(msg)
	//将文件作为附件返回r给客户端进行下载
	//调用函数传输文件
	c.FileAttachment(filePath, fileName)

	// 更新下载次数
	service.DownloadNumAdd(fileId)

	for i := 0; i < maxRetry; i++ {
		msg, err := global.ServiceSetup.LogAction(userID, "下载文件", fileName)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

}

func (FileApi) FileDownloadByIdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fmt.Println(user.IP)

	fileId := c.GetHeader("file_id")
	fileModel := models.FileModel{}
	fileModel = service.GetFileInfo(fileId)
	filePath := path.Join(global.Path, fileModel.FilePath, fileModel.FileName+fileModel.Postfix)
	fileName := fileModel.FileName + fileModel.Postfix
	fmt.Println(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		c.Status(500)
		return
	}
	defer file.Close()
	fmt.Println("文件打开成功")

	fileContent, _ := os.ReadFile(filePath)
	privateKeyPEM, err := os.ReadFile("private.pem")
	if err != nil {
		global.Log.Fatal(err)
		res.FailWithMessage("读取密钥失败", c)
		return
	}

	privateKey, err := valid.DecodePrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		global.Log.Fatal(err)
		res.FailWithMessage("读取密钥失败", c)
		return
	}
	decryptedFileContent, err := valid.DecryptWithRSA(privateKey, fileContent)
	//对比链上连下数据哈希是否相等
	maxRetry := 5
	var msg string

	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.VerifyDataHash(fileId, string(decryptedFileContent))
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
	}
	if msg != "File integrity check passed." {
		err = errors.New("链上连下数据协同验证不统一，下载请求被拒绝")
		c.Status(550)
		return
	}
	//将文件作为附件返回r给客户端进行下载
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "download-*.tmp")
	if err != nil {
		// 处理创建临时文件错误
		res.FailWithMessage("无法创建临时文件", c)
		return
	}
	defer os.Remove(tmpFile.Name())

	// 将解密后的文件内容写入临时文件
	_, err = tmpFile.Write(decryptedFileContent)
	if err != nil {
		// 处理写入文件错误
		res.FailWithMessage("无法写入文件", c)
		return
	}

	// 获取临时文件的路径和文件名
	filePathTemp := tmpFile.Name()

	//调用函数传输文件
	c.FileAttachment(filePathTemp, fileName)

	// 更新下载次数
	service.DownloadNumAdd(fileId)

	for i := 0; i < maxRetry; i++ {
		msg, err := global.ServiceSetup.LogAction(userID, "下载文件", fileName)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

}
