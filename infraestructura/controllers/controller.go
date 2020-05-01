package controllers

import "github.com/ant0ine/go-json-rest/rest"

type Controller interface {
	Routes() []*rest.Route
}
