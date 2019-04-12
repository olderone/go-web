package routers

import (
	"web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})

	ParkController := &controllers.ParkController{}
	beego.Router("/api/v1/hospList", ParkController, "post:GetParkList")
	beego.Router("/api/v1/hospInstallPack", ParkController, "post:HospInstallPack")
	beego.Router("/admin/park/GetOnePark", ParkController, "post:GetOnePark")

	ExtraController := &controllers.ExtraController{}
	beego.Router("/api/v1/hospExtraImg", ExtraController, "post:HospExtraImg")

	AddrController := &controllers.AddrController{}
	beego.Router("/api/v1/getCityList", AddrController, "post:GetCityList")
}
