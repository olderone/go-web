package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (admin *AdminController) Get() {
	admin.Data["WebName"] = "光享科技后台管理"
	admin.TplName = "admin/index.html"
}
