package evento_repository

import (
	"context"
	"time"

	"../../database"
	m "../../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Llamando a coleccion
var collection = database.GetCollection("Eventos")

//llamando contexto
var ctx = context.Background()

//Funcion de crear evento
func Create(evento m.Evento) error {
	var err error
	//Insertando en coleccion
	_, err = collection.InsertOne(ctx, evento)

	if err != nil {
		return err
	}
	return nil
}

///Funcion para mostrar la coleccion eventos
func Read() (m.Eventos, error) {
	var eventos m.Eventos
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var evento m.Evento
		err = cur.Decode(&evento)

		if err != nil {
			return nil, err
		}
		eventos = append(eventos, &evento)
	}

	return eventos, nil
}

func Update(evento m.Evento, idEvento string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(idEvento)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"descripcion": evento.Descripcion,
			"updated_at":  time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(IdEvento string) error {
	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(IdEvento)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
