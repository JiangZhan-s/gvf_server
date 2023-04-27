package user_ser

import (
	"gvf_server/service/redis_ser"
	"gvf_server/utils/jwts"
	"time"
)

type UserService struct {
}

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
