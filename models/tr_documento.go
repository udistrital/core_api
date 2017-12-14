package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrDocumento struct {
	Documento              *Documento
	ValorAtributoDocumento []*ValorAtributoDocumento
}

/*
	Función para la transaccion de solicitudes de pagos mensuales
*/
func AddDocumentos(m *TrDocumento) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	var id int64

	//Seteo fecha de creación transacción
	//m.Documento.FechaDocumento = time.Now() // se elimino este campo
	//m.PagoMensual.Mes = float64(m.PagoMensualEstadoPagoMensual.Fecha.Month())

	/*Inserta documento e inserta valoresAtributoDocumento*/
	if id, err = o.Insert(m.Documento); err == nil {
		fmt.Println("Atributo ", m.ValorAtributoDocumento)
		for _, v := range m.ValorAtributoDocumento {
			v.Documento = &Documento{Id: int(id)}
			//---
			if _, err = o.Insert(v); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los valores de los atributos del documento!")
				return
			}
		}
		o.Commit()
		alerta = append(alerta, "El documento ha sido insertado")
		return
	} else {
		o.Rollback()
		alerta[0] = "error"
		alerta = append(alerta, "Error: ¡Ocurrió un error al insertar el documento!")
		return
	}
	return alerta, err
}
