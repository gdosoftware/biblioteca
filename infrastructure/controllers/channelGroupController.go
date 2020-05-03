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
		rest.Get("/channelgroup/#id", http.RetrieveChannelGroupHttp),
		rest.Get("/channelgroup/application/#app", http.FindAllChannelGroupHttp),
		rest.Get("/channelgroup/application/#app/#type", http.FindAllChannelGroupByTypeHttp),

		rest.Post("/channelgroup", http.CreateChannelGroupHttp),
		rest.Put("/channelgroup/#id", http.UpdateChannelGroupHttp),
		rest.Delete("/channelgroup/#id", http.DeleteChannelGroupHttp),
	}
	controller.routes = routes
	return &controller
}

func (s *ChannelGroupController) Routes() []*rest.Route {
	return s.routes
}



