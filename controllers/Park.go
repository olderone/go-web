package controllers

import (
	"encoding/json"
	"web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ParkController struct {
	beego.Controller
}

//返回全部数据
func (c *ParkController) GetParkList() {

	dataList, err := models.QueryAllParkInfo()
	if err == nil {
		c.Data["json"] = dataList
		c.ServeJSON()
	}
	logs.Info("dataList :", dataList)

}

//查询单条数据
func (this *ParkController) GetOnePark() {
	//获得id
	id, _ := this.GetInt32("Id", 1)
	ParkOne, err := models.GetParkById(id)
	if err == nil {
		json.Unmarshal(this.Ctx.Input.RequestBody, &ParkOne)
		this.Data["json"] = ParkOne
		this.ServeJSON()
	} else {
		tmpHospInfo := &models.HospInfo{}

		this.Data["json"] = tmpHospInfo
		this.ServeJSON()
	}

}

// 版本控制
func (this *ParkController) HospInstallPack() {
	appver := this.GetString("appver")
	insPak, err := models.GetInstallPak(appver)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 1, "msg": "没有数据", "data": err}
	} else {

		this.Data["json"] = map[string]interface{}{"status": 0, "msg": "没有数据", "data": insPak}
	}

	this.ServeJSON()
}
