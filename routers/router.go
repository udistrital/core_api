package routers

import (
	"github.com/udistrital/crud_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
