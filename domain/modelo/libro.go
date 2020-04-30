package modelo

import "gopkg.in/mgo.v2/bson"

type Libro struct {
	id        bson.ObjectId `bson:"_id"`
	titulo    string        `bson:"titulo"`
	autor     string        `bson:"autor"`
	stock     int           `bson:"stock"`
	prestados int           `bson:"prestados"`
}
