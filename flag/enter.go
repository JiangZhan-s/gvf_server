package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string //-u admin
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止 web 项目
func IsWebStop(option Option) bool {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				return true
			}
		case bool:
			if val {
				return true
			}
		}
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
	}
	sys_flag.Usage()
}
