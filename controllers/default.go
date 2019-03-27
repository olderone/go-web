package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["WebName"] = "成都光享科技有限公司"
	c.Data["Website"] = "www.liteshare.cn"
	c.Data["Email"] = "business@liteshare.cn"
	//c.TplName = "index.tpl"
	c.TplName = "index.html"
}
