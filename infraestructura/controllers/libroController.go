package controller

type LibroController struct {
	ChannelGroupAPI *api.LibroApi
	routes          []*rest.Route
}

func CreateLibroController(api *api.LibroApi) *LibroController {
	controller := LibroController{LibroApi: api}
	routes := []*rest.Route{
		rest.Get("/libro/#id", api.GetOne),
		rest.Get("/libro/application/:app", api.GetAll),
		rest.Get("/libro/application/:app/:type", api.GetByType),

		rest.Post("/libro", api.Save),
		rest.Put("/libro/#id", api.Update),
		rest.Delete("/libro/#id", api.Delete),
	}
	controller.routes = routes
	return &controller
}

func (s *LibroController) Routes() []*rest.Route {
	return s.routes
}



