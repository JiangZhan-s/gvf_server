package file_store_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvf_server/models/res"
	"gvf_server/service"
	"gvf_server/utils/jwts"
)

func (FileStoreApi) StoreQueryByUserIdView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID

	//获取用户信息
	_, err := service.GetUserInfo(userID)
	if err != nil {
		res.FailWithMessage(fmt.Sprintf("未找到用户:%d", userID), c)
		return
	}

	storeInfo, _ := service.GetUserFileStore(int(userID))
	res.OkWithData(storeInfo, c)

}
