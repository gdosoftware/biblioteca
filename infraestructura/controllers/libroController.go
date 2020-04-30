package controller

import("github.com/gdosoftware/biblioteca/infraestructura/api")

type LibroController struct {
	LibroApi *api.LibroApi
	routes          []*rest.Route
}

func CreateLibroController(api *api.LibroApi) *LibroController {
	controller := LibroController{LibroApi: api}
	routes := []*rest.Route{
		rest.Get("/libro/#id", api.RecuperarLibro),
		rest.Get("/libro/", api.RecuperarTodosLosLibro),

		rest.Post("/libro", api.AltaLibro),
		rest.Put("/libro/#id", api.ModificacionLibro),
		rest.Delete("/libro/#id", api.BorrarLibro),
	}
	controller.routes = routes
	return &controller
}

func (s *LibroController) Routes() []*rest.Route {
	return s.routes
}



