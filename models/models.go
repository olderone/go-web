package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	//注册 model
	orm.RegisterModel(new(HospInfo), new(AdminUsers))

}
