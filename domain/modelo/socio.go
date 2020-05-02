package modelo

import "gopkg.in/mgo.v2/bson"

type Socio struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Dni     string        `bson:"dni"`
	Deleted   string      `bson:"deleted"`
}
