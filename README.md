## datereminder 轻量年度日程提醒小软件
一个简单日程提醒小程序，用于提醒用户年度任务例如亲友生日、车险续费、医保续费等等，支持公历、农历日期提醒，提醒方式为发送邮件给用户。
建议使用方式是邮件绑定到微信上，即可微信查看。

- 提醒示例:
![](https://github.com/del-xiong/datereminder/blob/master/demo.jpg?raw=true)

- 用法:
   1. 导入task.sql到数据库
   2. 修改config/param.go smtp登录信息、数据库链接信息 (基本不改动直接放代码了)
   3. 添加提醒任务数据，task_name/task_desc任务名和描述。is_lunar 任务日期模式 默认0公历 1为农历。is_loop 是否循环提醒(暂时没有使用全部为循环提醒)。pre_day 提前多少天开始发送提醒通知。
   4. 添加计划任务每日执行即可。
