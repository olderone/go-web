package controllers

import (
	"web/models"

	"github.com/astaxie/beego"
)

type AddrController struct {
	beego.Controller
}

func (this *AddrController) GetCityList() {

	city, err, num := models.GetCityList()

	for index := range city {
		city[index].CityName = "ChengDu"
		city[index].ProvinceName = "SiChuan"
		city[index].CountryName = "ZhongGuo"
	}

	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 1, "msg": "没有数据", "data": err}
	} else {

		this.Data["json"] = map[string]interface{}{"status": 0, "msg": "有数据", "data": city, "num": num}
	}

	this.ServeJSON()
}
