package routers

import (
	"github.com/udistrital/crud_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/tipo_entidad",
			beego.NSInclude(
				&controllers.Tipo_entidadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
