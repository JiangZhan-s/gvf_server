package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 //管理员
	SignTel   SignStatus = 2 //用户
	SignEmail SignStatus = 3 //游客
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignTel:
		str = "手机号"
	case SignEmail:
		str = "邮箱"
	default:
		str = "其他"
	}
	return str
}
