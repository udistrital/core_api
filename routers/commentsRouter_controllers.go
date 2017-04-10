package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:Tipo_entidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
			beego.ControllerComments{
				Method: "Post",
				Router: `/`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:id`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
			beego.ControllerComments{
				Method: "Put",
				Router: `/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:CiiuTipoController"],
			beego.ControllerComments{
				Method: "Delete",
				Router: `/:id`,
				AllowHTTPMethods: []string{"delete"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
			beego.ControllerComments{
				Method: "Post",
				Router: `/`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:id`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
			beego.ControllerComments{
				Method: "Put",
				Router: `/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:JefeDependenciaController"],
			beego.ControllerComments{
				Method: "Delete",
				Router: `/:id`,
				AllowHTTPMethods: []string{"delete"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
			beego.ControllerComments{
				Method: "Post",
				Router: `/`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:id`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
			beego.ControllerComments{
				Method: "Put",
				Router: `/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:OrdenadorGastoController"],
			beego.ControllerComments{
				Method: "Delete",
				Router: `/:id`,
				AllowHTTPMethods: []string{"delete"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "],
			beego.ControllerComments{
				Method: "Post",
				Router: `/`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:id`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "],
			beego.ControllerComments{
				Method: "Put",
				Router: `/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosDependenciaController                             "],
			beego.ControllerComments{
				Method: "Delete",
				Router: `/:id`,
				AllowHTTPMethods: []string{"delete"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
			beego.ControllerComments{
				Method: "Post",
				Router: `/`,
				AllowHTTPMethods: []string{"post"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: `/:id`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: `/`,
				AllowHTTPMethods: []string{"get"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
			beego.ControllerComments{
				Method: "Put",
				Router: `/:id`,
				AllowHTTPMethods: []string{"put"},
				Params: nil})

		beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/core_api/controllers:RubrosOrdenadorController"],
			beego.ControllerComments{
				Method: "Delete",
				Router: `/:id`,
				AllowHTTPMethods: []string{"delete"},
				Params: nil})
		}
