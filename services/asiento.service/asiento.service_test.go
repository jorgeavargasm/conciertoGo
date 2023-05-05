package asientoservice_test

import (
	m "../../models"
	AsientoService "../asiento.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var asientoId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	asientoId = oid.Hex()

	asiento := m.Asiento{
		Id:          oid,
		IdEvento:    "",
		Descripcion: "Ejemplo",
		Asiento:     10,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	err := AsientoService.Create(asiento)
	if err != nil {
		t.Error("la prueba de persistencia de datos de asientos a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}
func TestRead(t *testing.T) {
	asientos, err := AsientoService.Read()
	if err != nil {
		t.Error("Se ha presentado un error en la consulta del usuario")
		t.Fail()
	}
	if len(asientos) == 0 {
		t.Error("No se encontraron resultados")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}

}
func TestUpdate(t *testing.T) {
	asiento := m.Asiento{
		Descripcion: "Ejemplo update black",
	}

	err := AsientoService.Update(asiento, asientoId)

	if err != nil {
		t.Error("Error al tratar de actualizar el evento")
		t.Fail()
	} else {
		t.Log("La prueba de actualización finalizo con exito!")
	}
}
func TestDelete(t *testing.T) {

	err := AsientoService.Delete(asientoId)
	if err != nil {
		t.Error("Error al tratar de eliminar el evento")
		t.Fail()
	} else {
		t.Log("La prueba de eliminación finalizo con exito!")
	}
}
