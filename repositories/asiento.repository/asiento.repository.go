package asiento_repository

import (
	"context"
	"time"

	"../../database"
	m "../../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Llamando a coleccion
var collection = database.GetCollection("Asientos")

//llamando contexto
var ctx = context.Background()

//Funcion de crear asientos
func Create(asiento m.Asiento) error {
	var err error
	//Insertando en coleccion
	_, err = collection.InsertOne(ctx, asiento)

	if err != nil {
		return err
	}
	return nil
}

///Funcion para mostrar la coleccion asientos
func Read() (m.Asientos, error) {
	var asientos m.Asientos
	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var asiento m.Asiento
		err = cur.Decode(&asiento)

		if err != nil {
			return nil, err
		}
		asientos = append(asientos, &asiento)
	}

	return asientos, nil
}

func Update(asiento m.Asiento, idAsiento string) error {
	var err error
	oid, _ := primitive.ObjectIDFromHex(idAsiento)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"descripcion": asiento.Descripcion,
			"Asiento":     asiento.Asiento,
			"updated_at":  time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(idAsiento string) error {
	var err error
	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(idAsiento)

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
