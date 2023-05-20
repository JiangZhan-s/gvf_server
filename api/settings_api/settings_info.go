package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/global"
	"time"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//res.FailWithCode(2, c)
	maxRetry := 5 // 设置最大重试次数
	for i := 0; i < maxRetry; i++ {
		msg, err := global.ServiceSetup.QueryLedger()
		if err != nil {
			fmt.Printf("Error: %s, retrying...\n", err.Error())
		} else {
			fmt.Println(msg)
			break // 成功获取到结果，跳出循环
		}

		time.Sleep(1 * time.Second) // 暂停1秒后重试
	}

}
