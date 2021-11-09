package models

import (
	"time"
)

type Task struct {
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	IsLoop     int       `xorm:"not null default 0 comment('是否循环提醒') index INT"`
	IsLunar    int       `xorm:"not null default 0 comment('是否农历日
默认公历') INT"`
	PreDay       int       `xorm:"not null default 15 comment('提前多少天开始提醒') index INT"`
	ReminderDate string    `xorm:"default '' comment('提醒日期') VARCHAR(100)"`
	TaskDesc     string    `xorm:"not null default '' comment('描述') VARCHAR(500)"`
	TaskId       int       `xorm:"not null pk autoincr INT"`
	TaskName     string    `xorm:"not null default '' comment('任务名') VARCHAR(200)"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
