package src

import (
	"datereminder/config"
	"fmt"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"github.com/tidwall/gjson"
	gomail "gopkg.in/gomail.v2"
	"log"
	"strconv"
	"time"
)

func NotifyMsg(subject, body string) {
	d := gomail.NewDialer(config.SMTP_SERVER, 25, config.SMTP_EMAIL, config.SMTP_PWD)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	log.Printf("%s:\n%s\n\n", subject, body)

	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTP_EMAIL)
	m.SetHeader("To", config.SMTP_RECEIPT)
	m.SetHeader("Subject", subject)
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	body += fmt.Sprintf("<br><br>检查时间(公历): %s", time.Unix(time.Now().Unix()+8*3600, 0).UTC().Format("2006-01-02 15:04:05"))
	solarDate := solar2lunar(time.Now())
	body += fmt.Sprintf("<br> 检查时间(农历): %s", time.Unix(solarDate.Unix()+8*3600, 0).UTC().Format("2006-01-02"))
	m.SetBody("text/html", body)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Printf("发送邮件通知失败: %s\n", subject+body)
		log.Println(err)
	}
}

// 公历日期转农历
func solar2lunar(rdate time.Time) time.Time {
	year, _ := strconv.Atoi(rdate.Format("2006"))
	month, _ := strconv.Atoi(rdate.Format("01"))
	day, _ := strconv.Atoi(rdate.Format("02"))
	c := calendar.BySolar(int64(year), int64(month), int64(day), 0, 0, 0)
	json, _ := c.ToJSON()
	t, _ := time.Parse("2006-01-02", gjson.GetBytes(json, "lunar.year").String()+"-"+fmt.Sprintf("%02d", gjson.GetBytes(json, "lunar.month").Int())+"-"+fmt.Sprintf("%02d", gjson.GetBytes(json, "lunar.day").Int()))
	return t
}
