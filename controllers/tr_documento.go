package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/core_api/models"
)

// oprations for Documento
type TrDocumentoController struct {
	beego.Controller
}

func (c *TrDocumentoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title TrDocumento
// @Description insert the TrDocumento
// @Param	body		body 	models.TrDocumento	true	"body for TrDocumento content"
// @Success 200 {object} msg
// @Failure 403 :id is not int
// @router / [post]
func (c *TrDocumentoController) Post() {
	var v models.TrDocumento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.AddDocumentos(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
		} else {
			c.Data["json"] = alerta
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}
