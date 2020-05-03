package model

import "gopkg.in/mgo.v2/bson"

type ChannelGroup struct {
	Id            bson.ObjectId `bson:"_id"`
	Application   string        `bson:"application"`
	Name          string        `bson:"name"`
	Type          string        `bson:"type"`
	Channels      []Channel     `bson:"channels"`
	Deleted       string        `bson:"deleted"`
}

