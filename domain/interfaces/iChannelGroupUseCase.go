package interfaces

import "gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/domain/model"

type IChannelGroupUseCase interface {
	CreateChannelGroup(channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	UpdateChannelGroup(id string, channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	RetrieveChannelGroup(id string) (model.ChannelGroup, error)
	DeleteChannelGroup(id string) error
	FindAllChannelGroup(app string) ([]model.ChannelGroup, error)
	FindAllChannelGroupByType(app string, tipo string) ([]model.ChannelGroup, error)
}

