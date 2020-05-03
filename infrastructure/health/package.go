package health

import "gopkg.in/mgo.v2"

type Sensor interface {
	Name() string
	Health() error
	Stats() interface{}
}

type RepoStats struct {
	BuildInfo   mgo.BuildInfo `json:"mongoSession"`
	LiveServers []string      `json:"serversAlive"`
}
