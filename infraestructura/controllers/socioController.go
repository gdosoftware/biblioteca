package controller

import("github.com/gdosoftware/biblioteca/infraestructura/api")

type SocioController struct {
	socioApi        *api.SocioApi
	routes          []*rest.Route
}

func CreateLibroController(api *api.SocioApi) *SocioController {
	controller := SocioController{SocioApi: api}
	routes := []*rest.Route{
		rest.Get("/socio/#id", api.RecuperarSocio),
		rest.Get("/socio/", api.RecuperarTodosLosSocios),

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
