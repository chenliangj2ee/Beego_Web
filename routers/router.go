package routers

import (
	"Beego_Web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/login/phone:phone", &controllers.UserController{},"get:Login")
}
