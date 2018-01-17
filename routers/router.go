// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/core_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_entidad",
			beego.NSInclude(
				&controllers.Tipo_entidadController{},
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

		beego.NSNamespace("/tr_documento",
			beego.NSInclude(
				&controllers.TrDocumentoController{},
			),
		),

		beego.NSNamespace("/actividad_economica",
			beego.NSInclude(
				&controllers.ActividadEconomicaController{},
			),
		),

		beego.NSNamespace("/dominio_tipo_documento",
			beego.NSInclude(
				&controllers.DominioTipoDocumentoController{},
			),
		),

		beego.NSNamespace("/clasificacion_ciiu",
			beego.NSInclude(
				&controllers.ClasificacionCiiuController{},
			),
		),

		beego.NSNamespace("/tipo_dato",
			beego.NSInclude(
				&controllers.TipoDatoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
