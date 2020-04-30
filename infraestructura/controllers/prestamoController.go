package controller

import("github.com/gdosoftware/biblioteca/infraestructura/api")

type PrestamoController struct {
	PrestamoApi *api.PrestamoApi
	routes          []*rest.Route
}

func CreateLibroController(api *api.PrestamoApi) *PrestamoController {
	controller := PrestamoController{PrestamoApi: api}
	routes := []*rest.Route{
			rest.Post("/prestamo/#libroId/#socioId", api.PrestarLibro),
		rest.Put("/prestamo/#prestamoId", api.DevolverLibro),
	}
	controller.routes = routes
	return &controller
}

func (s *PrestamoController) Routes() []*rest.Route {
	return s.routes
}
