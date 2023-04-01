package file_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils"
	"gvf_server/utils/jwts"
	"io"
	"net/http"
	"os"
)

func (FileApi) FileUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
	}
	Fid := c.GetHeader("id")
	//接收上传文件
	file, head, err := c.Request.FormFile("file")
	res.OkWithData(head, c)
	//判断当前文件夹是否有同名文件
	if ok := service.CurrFileExists(Fid, head.Filename); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
		})
		return
	}

	//判断用户的容量是否足够
	if ok := service.CapacityIsEnough(head.Size, user.FileStoreID); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 503,
		})
		return
	}

	if err != nil {
		fmt.Println("文件上传错误", err.Error())
		return
	}
	defer file.Close()

	//在本地创建一个新的文件
	newFile, err := os.Create(global.Path + "\\" + head.Filename) //+ "\\" + user.UserName
	if err != nil {
		fmt.Println(err)
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Println("文件拷贝错误", err.Error())
		return
	}

	//将光标移至开头
	_, _ = newFile.Seek(0, 0)
	utils.GetSHA256HashCode(newFile)

	//新建文件信息
	service.CreateFile(head.Filename, fileSize, Fid, user.FileStoreID, int(user.ID))
	//上传成功减去相应剩余容量
	service.SubtractSize(fileSize, user.FileStoreID)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
