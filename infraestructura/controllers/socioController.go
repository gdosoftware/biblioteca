package controllers

import("github.com/gdosoftware/biblioteca/infraestructura/api"
		"github.com/ant0ine/go-json-rest/rest")

type SocioController struct {
	SocioApi        *api.SocioApi
	routes          []*rest.Route
}

func CreateSocioController(api *api.SocioApi) *SocioController {
	controller := SocioController{SocioApi: api}
	routes := []*rest.Route{
		rest.Get("/socio/#id", api.RecuperarSocio),
		rest.Get("/socio", api.RecuperarTodosLosSocios),

		rest.Post("/socio", api.AltaSocio),
		rest.Put("/socio/#id", api.ModificacionSocio),
		rest.Delete("/socio/#id", api.BorrarSocio),
	}
	controller.routes = routes
	return &controller
}

func (s *SocioController) Routes() []*rest.Route {
	return s.routes
}
