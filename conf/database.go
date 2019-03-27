package conf

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func RegistDB() {
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/liteshare?charset=utf8") //密码为空格式

}
