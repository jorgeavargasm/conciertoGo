package models

import (
	"time"
)

type Asiento_individual struct {
	Descripcion string    `json:Descripcion`
	Asiento     int       `json:Asiento`
	CreateAt    time.Time `bson:"created_at" json:"created_at"`
	UpdateAt    time.Time `bson:"updated_at" json:"updated_at, omitempty"`
}

type Asiento_individuales []*Asiento_individual
