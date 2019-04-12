package controllers

import (
	"strconv"
	"web/models"

	"github.com/astaxie/beego"
)

type ExtraController struct {
	beego.Controller
}

func (this *ExtraController) HospExtraImg() {
	hospNo := this.Input().Get("hospNo")
	intid, err := strconv.Atoi(hospNo)

	extra, err, num := models.GetExtra(intid)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 1, "msg": "没有数据", "data": err}
	} else {

		this.Data["json"] = map[string]interface{}{"status": 0, "msg": "有数据", "data": extra, "num": num}
	}

	this.ServeJSON()
}
