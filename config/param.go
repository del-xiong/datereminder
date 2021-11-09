package config

import "github.com/go-xorm/xorm"

var (
	SMTP_SERVER  = "smtpdm.aliyun.com"
	SMTP_EMAIL   = "test2@alimail.alidddd.com" // smtp账号
	SMTP_PWD     = "aaaaaaaaaaaaaaa"           // smtp密码
	SMTP_RECEIPT = "me@qq.com"                 // 接收通知人
	MysqlConnect = "datereminder:password@tcp(localhost)/datereminder?charset=utf8mb4"
)

var Engine *xorm.Engine
