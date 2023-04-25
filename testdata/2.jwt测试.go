package main

import (
	"fmt"
	"gvf_server/core"
	"gvf_server/global"
	"gvf_server/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		NickName: "xxx",
	})
	fmt.Println(token, err)
	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImppYW5nemhhbiIsIm5pY2tfbmFtZSI6Inh4eCIsInJvbGUiOjEsInVzZXJfaWQiOjEsImV4cCI6MTY3OTU3MjQ0MC4xMDM2ODMsImlzcyI6IjEyMzQifQ.MQ-o5fGRrfTIbUjPFk_jSeil6Y5P3MLmgxCr4fdfwXM")
	fmt.Println(claims, err)
}
