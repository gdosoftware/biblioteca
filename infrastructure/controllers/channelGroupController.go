package controllers

import("github.com/gdosoftware/biblioteca/infrastructure/http"
		"github.com/ant0ine/go-json-rest/rest")

type ChannelGroupController struct {
	ChannelGroupHttp       *http.ChannelGroupHttp
	routes          []*rest.Route
}

func CreateChannelGroupController(http *http.ChannelGroupHttp) *ChannelGroupController {
	controller := ChannelGroupController{ChannelGroupHttp: http}
	routes := []*rest.Route{
		rest.Get("/channelGroup/#id", http.RecuperarChannelGroup),
		rest.Get("/channelGroup", http.RecuperarTodosLosChannelGroups),

		rest.Post("/channelGroup", http.AltaChannelGroup),
		rest.Put("/channelGroup/#id", http.ModificacionChannelGroup),
		rest.Delete("/channelGroup/#id", http.BorrarChannelGroup),
	}
	controller.routes = routes
	return &controller
}

func (s *ChannelGroupController) Routes() []*rest.Route {
	return s.routes
}



