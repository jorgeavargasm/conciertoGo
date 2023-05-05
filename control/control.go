package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	m "../models"

	AsientoService "../services/asiento.service"
	EventoService "../services/evento.service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var eventoId string

func Controlador() {
	fmt.Println("Ejecutando control")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/listar/", listar).Methods("GET")
	router.HandleFunc("/guardar/", guardar).Methods("POST")
	fmt.Println("Servidor corriendo")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a apirest")
}
func listar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventos, err := EventoService.Read()
	if err != nil {
		respuesta := map[string]bool{"result": false}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	json.NewEncoder(w).Encode(eventos)
}

func guardar(w http.ResponseWriter, r *http.Request) {
	var asiento2 respuestaEventoDto
	arr := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	w.Header().Set("Content-Type", "application/json")

	var NuevoR eventoDto
	var respuestaG respuestaEventoDto
	//Obteniendo reqBody de la peticion
	reqBody, err := ioutil.ReadAll(r.Body)

	//Apuntado valores del ReqBody al modelo eventoDto
	json.Unmarshal(reqBody, &NuevoR)

	if err != nil {
		respuesta := map[string]bool{"result": false}
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	//Generando Id
	oid := primitive.NewObjectID()
	eventoId = oid.Hex()

	//Contruyendo objeto antes de enviarlo a mongo
	evento := m.Evento{
		Id:          oid,
		Descripcion: NuevoR.NombreSala,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	errx := EventoService.Create(evento)
	respuestaG.NombreSala = NuevoR.NombreSala
	respuestaG.CantidadFilas = NuevoR.CantidadFilas
	respuestaG.CantidadAsientosF = NuevoR.CantidadAsientosF

	var contador = 0
	var acumMo = 1
	for i := 0; i < NuevoR.CantidadFilas; i++ {
		var asiento_individual Asiento_individual2
		contador++
		if contador > 26 {
			acumMo++
			contador = 1
		}
		descArr := ""
		for j := 0; j < acumMo; j++ {
			descArr = descArr + arr[contador-1]
		}

		oid2 := primitive.NewObjectID()

		asiento2.NombreSala = NuevoR.NombreSala
		asiento2.CantidadFilas = NuevoR.CantidadFilas
		asiento2.CantidadAsientosF = NuevoR.CantidadAsientosF

		for k := 0; k < NuevoR.CantidadAsientosF; k++ {
			asiento_individual = Asiento_individual2{
				Descripcion: descArr,
				Asiento:     k,
			}
			asiento2.addArray(asiento_individual)
		}

		asiento := m.Asiento{
			Id:          oid2,
			IdEvento:    eventoId,
			Descripcion: descArr,
			Asiento:     NuevoR.CantidadAsientosF,
			CreateAt:    time.Now(),
			UpdateAt:    time.Now(),
		}

		erry := AsientoService.Create(asiento)
		if erry != nil {
			respuesta := map[string]bool{"result": false}
			json.NewEncoder(w).Encode(respuesta)
			return
		}
	}

	if errx != nil {
		respuesta := map[string]bool{"result": false}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	if errx != nil {

		respuesta := map[string]bool{"result": false}
		json.NewEncoder(w).Encode(respuesta)
		return
	}

	//respuesta := map[string]bool{"result": true}
	json.NewEncoder(w).Encode(asiento2)

}

type eventoDto struct {
	NombreSala        string `json:Descripcion`
	CantidadFilas     int    `json:CantidadFilas`
	CantidadAsientosF int    `json:CantidadAsientosF`
}
type respuestaEventoDto struct {
	NombreSala        string
	CantidadFilas     int
	CantidadAsientosF int
	Asientos          []Asiento_individual2
}

type Asiento_individual2 struct {
	Descripcion string `json:Descripcion`
	Asiento     int    `json:Asiento`
}

func (a *respuestaEventoDto) addArray(as Asiento_individual2) {

	a.Asientos = append(a.Asientos, as)
}
