package interfaces

import "gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/domain/model"

type IChannelGroupRepository interface {
	Create(channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	Update(id string, channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	Retrieve(id string) (model.ChannelGroup, error)
	Delete(id string) error
	
	FindAll(app string) ([]model.ChannelGroup, error)
	FindByType(app string, tipo string) ([]model.ChannelGroup, error)
}
