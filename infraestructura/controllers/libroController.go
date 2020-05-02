package controllers

import("github.com/gdosoftware/biblioteca/infraestructura/http"
		"github.com/ant0ine/go-json-rest/rest")

type LibroController struct {
	LibroHttp       *http.LibroHttp
	routes          []*rest.Route
}

func CreateLibroController(http *http.LibroHttp) *LibroController {
	controller := LibroController{LibroHttp: http}
	routes := []*rest.Route{
		rest.Get("/libro/#id", http.RecuperarLibro),
		rest.Get("/libro", http.RecuperarTodosLosLibros),

		rest.Post("/libro", http.AltaLibro),
		rest.Put("/libro/#id", http.ModificacionLibro),
		rest.Delete("/libro/#id", http.BorrarLibro),
	}
	controller.routes = routes
	return &controller
}

func (s *LibroController) Routes() []*rest.Route {
	return s.routes
}



