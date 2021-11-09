package src

import (
	"datereminder/models"
	"fmt"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

// 返回还有多少天到期以及到期日的公历日期
func CheckDate(task models.Task, reminderData string) (int, time.Time, error) {
	rdate, err := time.Parse("2006-01-02 15:04:05", reminderData)
	if err != nil {
		return 0, time.Time{}, err
	}
	// 农历
	if task.IsLunar == 1 {
		// 转公历计算
		rdate = lunar2solar(rdate)
	}

	if rdate.Unix() < time.Now().Unix() {
		return -1, rdate, nil
	}

	return int(float64(rdate.Unix()-time.Now().Unix())/86400 + 0.5), rdate, nil
}

// 农历日期转公历
func lunar2solar(rdate time.Time) time.Time {
	year, _ := strconv.Atoi(rdate.Format("2006"))
	month, _ := strconv.Atoi(rdate.Format("01"))
	day, _ := strconv.Atoi(rdate.Format("02"))
	c := calendar.ByLunar(int64(year), int64(month), int64(day), 0, 0, 0, false)
	json, _ := c.ToJSON()
	t, _ := time.Parse("2006-01-02", gjson.GetBytes(json, "solar.year").String()+"-"+fmt.Sprintf("%02d", gjson.GetBytes(json, "solar.month").Int())+"-"+fmt.Sprintf("%02d", gjson.GetBytes(json, "solar.day").Int()))
	return t
}
