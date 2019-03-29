package controllers

import (
	"fmt"
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
		c.Data["List"] = dataList
		c.Data["json"] = "添加成功"
		fmt.Println("Mark changed his speciality")
		c.ServeJSON()
	}
	logs.Info("dataList :", dataList)

}
