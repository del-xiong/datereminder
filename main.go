package main

import (
	"datereminder/config"
	"datereminder/models"
	"datereminder/src"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"strconv"
	"time"
)

type notifymsg struct {
	taskName string
	expDay   int
	content  string
}

func init() {
	flag.StringVar(&config.SMTP_SERVER, "smtp_server", "", "smtp服务器")
	flag.StringVar(&config.SMTP_EMAIL, "smtp_user", "", "smtp邮件登录账号")
	flag.StringVar(&config.SMTP_PWD, "smtp_pwd", "", "smtp邮件登录密码")
	flag.StringVar(&config.SMTP_RECEIPT, "smtp_receipt", "", "邮件通知收件人")
	flag.StringVar(&config.MysqlConnect, "mysql", "", "mysql登录信息 user:pwd@tcp(localhost)/datereminder")
	flag.Parse()

	Timeloc, _ := time.LoadLocation("PRC")
	var err error
	config.Engine, err = xorm.NewEngine("mysql", config.MysqlConnect)
	if err != nil {
		log.Fatalln(err)
	}
	config.Engine.DatabaseTZ = Timeloc
	config.Engine.TZLocation = Timeloc
}

func main() {
	msg := make([]notifymsg, 0)

	// 读取全部数据
	tasks := make([]models.Task, 0)
	handlerErr(config.Engine.Find(&tasks))

	if len(tasks) == 0 {
		// 没有任何计划 直接报异常
		src.NotifyMsg(
			"查询数据库失败或无任何计划 datereminder",
			"查询数据库失败或无任何计划 datereminder",
		)
		return
	}

	for _, v := range tasks {
		year, _ := strconv.Atoi(time.Now().Format("2006"))
		for _, d := range []string{fmt.Sprintf("%d-%s 00:00:00", year, v.ReminderDate), fmt.Sprintf("%d-%s 00:00:00", year+1, v.ReminderDate)} {
			dayExp, dateExp, err := src.CheckDate(v, d)
			if err != nil {
				src.NotifyMsg(
					fmt.Sprintf("%s 检测失败，请检查参数是否正确", v.TaskName),
					fmt.Sprintf("%s 检测失败，请检查参数是否正确", v.TaskName),
				)
				continue
			}
			// 已到达可提醒日期 开始发送提醒通知
			if dayExp <= v.PreDay && dayExp >= 0 {
				var datetips string = fmt.Sprintf("到期日<b>公历</b>%s", v.ReminderDate)
				if v.IsLunar == 1 {
					datetips = fmt.Sprintf("到期日<b style=\"color:red;\">农历</b>%s (公历%s)", v.ReminderDate, dateExp.Format("2006-01-02"))
				}
				msg = append(msg, notifymsg{
					taskName: v.TaskName,
					expDay:   dayExp,
					content:  fmt.Sprintf("%s 剩余%d天 %s<br><span style=\"color:#444;font-size:12px;\">%s</span>", v.TaskName, dayExp, datetips, v.TaskDesc),
				})
			}
		}
	}

	// 全部任务检查完毕 发送

	if len(msg) == 0 {
		src.NotifyMsg(
			"近期没有即将到期的任务 datereminder",
			"近期没有即将到期的任务 datereminder",
		)
	} else {
		content := ""
		for _, v := range msg {
			content += fmt.Sprintf("<b>%s</b><br>%s<br><br>", v.taskName, v.content)
		}
		src.NotifyMsg(
			fmt.Sprintf("【重要】%s(%d天)等%d个任务即将到期", msg[0].taskName, msg[0].expDay, len(msg)),
			content,
		)
	}
}

func handlerErr(err error) {
	if err != nil {
		log.Output(2, err.Error())
	}
}
