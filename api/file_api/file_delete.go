package file_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
	"os"
	"strconv"
)

func (FileApi) FileDeleteView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	fileId, err := strconv.Atoi(c.GetHeader("file_id"))

	file := service.GetFileInfo(c.GetHeader("file_id"))
	if err != nil {
		res.FailWithMessage("获取文件信息失败", c)
		global.Log.Error(err)
		return
	}
	filePath := global.Path + file.FilePath + "/" + file.FileName + file.Postfix
	err = os.Remove(filePath)
	if err != nil {
		res.FailWithMessage("删除文件失败", c)
		global.Log.Error(err)
		return
	}

	err = service.DeleteFileById(fileId)
	if err != nil {
		res.FailWithMessage("失败了", c)
		global.Log.Error(err)
		return
	}
	res.OkWithMessage("删除成功", c)
}
