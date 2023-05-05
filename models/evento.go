package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Evento struct {
	Id          primitive.ObjectID `bson:"_id, omitempty" json:"Id, omitempty"`
	Descripcion string             `json:"descripcion"`
	CreateAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at, omitempty"`
}
type Eventos []*Evento
