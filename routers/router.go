// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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

		beego.NSNamespace("/ciiu_tipo",
			beego.NSInclude(
				&controllers.CiiuTipoController{},
			),
		),

		beego.NSNamespace("/jefe_dependencia",
			beego.NSInclude(
				&controllers.JefeDependenciaController{},
			),
		),

		beego.NSNamespace("/ordenador_gasto",
			beego.NSInclude(
				&controllers.OrdenadorGastoController{},
			),
		),

		beego.NSNamespace("/rubros_dependencia",
			beego.NSInclude(
				&controllers.RubrosDependenciaController{},
			),
		),

		beego.NSNamespace("/rubros_ordenador",
			beego.NSInclude(
				&controllers.RubrosOrdenadorController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
