package file_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils"
	"gvf_server/utils/jwts"
	"io"
	"os"
	"time"
)

func (FileApi) FileUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fmt.Println(userID)
	folderID := c.GetHeader("folder_id")
	fmt.Println(folderID)
	//接收上传文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		res.FailWithMessage("文件上传错误", c)
		res.FailWithError(err, folderID, c)
		return
	}
	defer file.Close()
	//判断当前文件夹是否有同名文件
	if service.CurrFileExists(folderID, header.Filename) {
		res.FailWithMessage("当前文件夹已有同名文件存在", c)
		return
	}

	//判断用户的容量是否足够
	if !service.CapacityIsEnough(header.Size, int(user.ID)) {
		res.FailWithMessage("用户容量不足", c)
		return
	}

	var fileFolder models.FileFolderModel
	global.DB.Find(&fileFolder, "id = ?", folderID)
	folderPath := service.GetCurrentFolderPath(fileFolder)
	fmt.Println(folderPath)

	newFile, err := os.Create(global.Path + "/" + folderPath + "/" + header.Filename)
	if err != nil {
		res.FailWithMessage("文件创建失败", c)
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		res.FailWithMessage("文件拷贝错误", c)
		return
	}

	//将光标移至开头
	_, err = newFile.Seek(0, 0)
	if err != nil {
		res.FailWithMessage("文件光标移动错误", c)
		return
	}
	hashData := utils.GetSHA256HashCode(newFile)
	fmt.Println(hashData)
	//新建文件信息
	fileID := service.CreateFile("/"+user.UserName, header.Filename, fileSize, folderID, user.FileStoreID, int(user.ID))
	fmt.Println(fileID)
	//上传成功减去相应剩余容量
	service.SubtractSize(fileSize, user.FileStoreID)

	maxRetry := 5 // 设置最大重试次数
	for i := 0; i < maxRetry; i++ {
		msg, err := global.ServiceSetup.StoreDataHash(fileID, hashData)
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}

		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

	res.OkWithMessage(fmt.Sprintf("文件%s上传成功", header.Filename), c)
}

func (FileApi) MultiFileUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}

	folderID := c.GetHeader("folder_id")

	// 接收上传文件
	form, err := c.MultipartForm()
	fmt.Println(form)
	if err != nil {
		res.FailWithMessage("文件上传错误", c)
		return
	}
	files := form.File["file"]
	fmt.Println(files)
	for _, file := range files {
		// 读取文件
		inFile, err := file.Open()
		if err != nil {
			res.FailWithMessage("文件打开错误", c)
			return
		}
		defer inFile.Close()

		// 创建新文件
		newFile, err := os.Create(global.Path + "/" + user.UserName + "/" + file.Filename)
		if err != nil {
			res.FailWithMessage("文件创建失败", c)
			return
		}
		defer newFile.Close()

		// 拷贝文件
		fileSize, err := io.Copy(newFile, inFile)
		if err != nil {
			res.FailWithMessage("文件拷贝错误", c)
			return
		}

		// 将光标移至开头
		_, err = newFile.Seek(0, 0)
		if err != nil {
			res.FailWithMessage("文件光标移动错误", c)
			return
		}

		// 计算文件哈希值
		hashData := utils.GetSHA256HashCode(newFile)

		// 新建文件信息
		fileID := service.CreateFile("/"+user.UserName, file.Filename, fileSize, folderID, user.FileStoreID, int(user.ID))

		// 上传成功减去相应剩余容量
		service.SubtractSize(fileSize, user.FileStoreID)

		// 存储文件哈希值
		maxRetry := 5 // 设置最大重试次数
		for i := 0; i < maxRetry; i++ {
			msg, err := global.ServiceSetup.StoreDataHash(fileID, hashData)
			if err != nil {
				fmt.Printf("Error: %s, retrying...\n", err.Error())
			} else {
				fmt.Println(msg)
				break // 成功获取到结果，跳出循环
			}

			time.Sleep(1 * time.Second) // 暂停1秒后重试
		}
		for i := 0; i < maxRetry; i++ {
			msg, err := global.ServiceSetup.QueryDataHash(fileID)
			if err != nil {
				fmt.Printf("Error: %s, retrying...\n", err.Error())
			} else {
				fmt.Println(msg)
				break // 成功获取到结果，跳出循环
			}

			time.Sleep(1 * time.Second) // 暂停1秒后重试
		}

		res.OkWithMessage(fmt.Sprintf("文件[%s]上传成功", file.Filename), c)
	}

	res.OkWithMessage("所有文件上传完成", c)
}
