package controllers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/infraestructura/http"
)

type SystemController struct {
	SystemStatusHttp *http.SystemStatusHttp
	routes       []*rest.Route
}

func NewSystemController(http *http.SystemStatusHttp) *SystemController {
	controller := SystemController{SystemStatusHttp: http}
	routes := []*rest.Route{
		rest.Get("/info", http.Info),
		rest.Get("/stats", http.Status),
		rest.Get("/stats/http", http.Stats),
		rest.Get("/live", http.Health),
		rest.Get("/ready", http.Health),
	}
	controller.routes = routes
	return &controller
}

func (s *SystemController) Routes() []*rest.Route {
	return s.routes
}
