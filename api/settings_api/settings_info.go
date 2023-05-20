package settings_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
	"time"
)

type Log struct {
	Timestamp string // 日志时间戳
	UserID    string // 用户ID
	Action    string // 操作类型
	Details   string // 操作详情
	// 其他日志属性
}

type Ledger struct {
	Key      string                 `json:"key"`
	DataType string                 `json:"dataType"`
	Data     map[string]interface{} `json:"data"`
}

func (SettingsApi) ActionInfoView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}
	if user.UserName != "root" {
		res.FailWithMessage("权限不足，无法查看", c)
		return
	}

	maxRetry := 5 // 设置最大重试次数
	var msg string
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.QueryLogs()
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

	var d []Log
	if msg != "null" {
		//解析json数据
		err = json.Unmarshal([]byte(msg), &d)
		if err != nil {
			panic(err)
		}
	}

	res.OkWithData(d, c)
}

func (SettingsApi) LedgerInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	user, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}

	if user.UserName != "root" {
		res.FailWithMessage("权限不足，无法查看", c)
		return
	}

	maxRetry := 5 // 设置最大重试次数
	var msg string
	for i := 0; i < maxRetry; i++ {
		msg, err = global.ServiceSetup.QueryLedger()
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}
		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

	var d []Ledger
	if msg != "null" {
		//解析json数据
		err = json.Unmarshal([]byte(msg), &d)
		if err != nil {
			panic(err)
		}
	}

	res.OkWithData(d, c)

}
