package jwts

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"gvf_server/global"
)

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
