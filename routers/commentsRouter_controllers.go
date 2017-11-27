package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:AtributoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:BancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuDivisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuSubclaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:PuntoSalarialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SalarioMinimoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesAreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SniesNucleoBasicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:SucursalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TrDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:TrDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:ValorAtributoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
