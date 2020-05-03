package module

import (
	"github.com/gdosoftware/biblioteca/infrastructure/http"
	"github.com/gdosoftware/biblioteca/infrastructure/controllers"
	"github.com/gdosoftware/biblioteca/infrastructure"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/usecase"
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



