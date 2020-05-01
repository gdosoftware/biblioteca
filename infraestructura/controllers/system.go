package controllers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/infraestructura/api"
)

type SystemController struct {
	sysStatusAPI *api.SysStatusAPI
	routes       []*rest.Route
}

func NewSystemController(api *api.SysStatusAPI) *SystemController {
	controller := SystemController{sysStatusAPI: api}
	routes := []*rest.Route{
		rest.Get("/info", api.Info),
		rest.Get("/stats", api.Status),
		rest.Get("/stats/http", api.Stats),
		rest.Get("/live", api.Health),
		rest.Get("/ready", api.Health),
	}
	controller.routes = routes
	return &controller
}

func (s *SystemController) Routes() []*rest.Route {
	return s.routes
}
