package modelo

import "gopkg.in/mgo.v2/bson"

type Libro struct {
	Id        bson.ObjectId `bson:"_id"`
	Titulo    string        `bson:"titulo"`
	Autor     string        `bson:"autor"`
	Stock     int           `bson:"stock"`
	Prestados int           `bson:"prestados"`
}
