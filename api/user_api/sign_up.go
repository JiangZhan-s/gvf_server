package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models"
	"gvf_server/models/ctype"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/pwd"
)

type SignUpRequest struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	NickName   string `json:"nick_name"`
}

// SignUpView 注册
func (UserApi) SignUpView(c *gin.Context) {
	var cr SignUpRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//判断用户名是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name= ? ", cr.UserName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名已存在，请重新输入")
		res.FailWithMessage("用户名已存在，请重新输入", c)
		return
	}

	//检验两次密码
	if cr.Password != cr.RePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		res.FailWithMessage("两次密码不一致，请重新输入", c)
		return
	}

	//对密码进行哈希
	hashPwd := pwd.HashPwd(cr.RePassword)

	role := ctype.PermissionUser

	user := models.UserModel{
		NickName:   cr.NickName,
		UserName:   cr.UserName,
		Password:   hashPwd,
		Email:      cr.Email,
		Role:       role,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}
	//入库
	err = global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	fileStore, err := service.CreateFileStore(int(user.ID))
	if err != nil {
		global.Log.Error(err)
		return
	}
	_, err = service.CreateFolderRoot(int(fileStore.ID), user.UserName)
	if err != nil {
		global.Log.Error(err)
		return
	}
	user.FileStoreModel = *fileStore
	user.FileStoreID = int(fileStore.ID)
	global.DB.Save(&user)
	global.Log.Info(fmt.Sprintf("用户%s创建成功", cr.UserName))
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功", cr.UserName), c)
}
