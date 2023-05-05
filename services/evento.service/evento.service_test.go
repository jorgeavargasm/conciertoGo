package evento_service_test

import (
	m "../../models"
	EventoService "../evento.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var eventoId string

func TestCreate(t *testing.T) {
	oid := primitive.NewObjectID()
	eventoId = oid.Hex()

	evento := m.Evento{
		Id:          oid,
		Descripcion: "Ejemplo",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	err := EventoService.Create(evento)
	if err != nil {
		t.Error("la prueba de persistencia de datos de eventos a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}
func TestRead(t *testing.T) {
	eventos, err := EventoService.Read()
	if err != nil {
		t.Error("Se ha presentado un error en la consulta del usuario")
		t.Fail()
	}
	if len(eventos) == 0 {
		t.Error("No se encontraron resultados")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}

}
func TestUpdate(t *testing.T) {
	evento := m.Evento{
		Descripcion: "Ejemplo update black",
	}

	//err := EventoService.Update(evento, "000000000000000000000000")
	err := EventoService.Update(evento, eventoId)

	if err != nil {
		t.Error("Error al tratar de actualizar el evento")
		t.Fail()
	} else {
		t.Log("La prueba de actualización finalizo con exito!")
	}
}
func TestDelete(t *testing.T) {

	//err := EventoService.Delete("63e6a943427652e2ba93f910")
	err := EventoService.Delete(eventoId)
	if err != nil {
		t.Error("Error al tratar de eliminar el evento")
		t.Fail()
	} else {
		t.Log("La prueba de eliminación finalizo con exito!")
	}
}
