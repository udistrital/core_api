// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/core_api/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/tipo_entidad",
			beego.NSInclude(
				&controllers.Tipo_entidadController{},
			),
		),

		beego.NSNamespace("/ciiu_division",
			beego.NSInclude(
				&controllers.CiiuDivisionController{},
			),
		),

		beego.NSNamespace("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),
		),

		beego.NSNamespace("/snies_area",
			beego.NSInclude(
				&controllers.SniesAreaController{},
			),
		),
		beego.NSNamespace("/snies_nucleo_basico",
			beego.NSInclude(
				&controllers.SniesNucleoBasicoController{},
			),
		),

		beego.NSNamespace("/ciiu_clase",
			beego.NSInclude(
				&controllers.CiiuClaseController{},
			),
		),

		beego.NSNamespace("/ciiu_subclase",
			beego.NSInclude(
				&controllers.CiiuSubclaseController{},
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

		beego.NSNamespace("/sucursal",
			beego.NSInclude(
				&controllers.SucursalController{},
			),
		),

		beego.NSNamespace("/banco",
			beego.NSInclude(
				&controllers.BancoController{},
			),
		),

		beego.NSNamespace("/punto_salarial",
			beego.NSInclude(
				&controllers.PuntoSalarialController{},
			),
		),

		beego.NSNamespace("/salario_minimo",
			beego.NSInclude(
				&controllers.SalarioMinimoController{},
			),
		),

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/atributo_documento",
			beego.NSInclude(
				&controllers.AtributoDocumentoController{},
			),
		),

		beego.NSNamespace("/valor_atributo_documento",
			beego.NSInclude(
				&controllers.ValorAtributoDocumentoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
