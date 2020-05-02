package controllers

import("github.com/gdosoftware/biblioteca/infraestructura/http"
		"github.com/ant0ine/go-json-rest/rest")

type SocioController struct {
	SocioHttp       *http.SocioHttp
	routes          []*rest.Route
}

func CreateSocioController(http *http.SocioHttp) *SocioController {
	controller := SocioController{SocioHttp: http}
	routes := []*rest.Route{
		rest.Get("/socio/#id", http.RecuperarSocio),
		rest.Get("/socio", http.RecuperarTodosLosSocios),

		rest.Post("/socio", http.AltaSocio),
		rest.Put("/socio/#id", http.ModificacionSocio),
		rest.Delete("/socio/#id", http.BorrarSocio),
	}
	controller.routes = routes
	return &controller
}

func (s *SocioController) Routes() []*rest.Route {
	return s.routes
}
