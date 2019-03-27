package routers

import (
	"web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})

	ParkController := &controllers.ParkController{}
	beego.Router("/admin/park/getParkList", ParkController, "post:GetParkList")
}
