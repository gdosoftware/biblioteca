package module

import (
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/infrastructure/http"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/infrastructure/controllers"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/infrastructure"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/domain/interfaces"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/domain/usecase"
)

func MakeControllers(
	iChannelGroupRepository interfaces.IChannelGroupRepository,
	
) []server.Controller {
	
	channelGroupController := controllers.CreateChannelGroupController(createChannelGroupHttp(iChannelGroupRepository))

	return []server.Controller{channelGroupController}
}

func createChannelGroupHttp(repo interfaces.IChannelGroupRepository) *http.ChannelGroupHttp{
   return http.CreateChannelGroupHttp(usecase.CreateChannelGroupUseCaseImpl(repo))
}



