package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asiento struct {
	Id          primitive.ObjectID `bson:"_id, omitempty" json:"Id, omitempty"`
	IdEvento    string             `json:IdEvento`
	Descripcion string             `json:Descripcion`
	Asiento     int                `json:Asiento`
	CreateAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time          `bson:"updated_at" json:"updated_at, omitempty"`
}

type Asientos []*Asiento
