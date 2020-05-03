package interfaces

import "github.com/gdosoftware/biblioteca/domain/model"

type IChannelGroupUseCase interface {
	CreateChannelGroup(channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	UpdateChannelGroup(id string, channelGroup *model.ChannelGroup) (*model.ChannelGroup, error)
	RetrieveChannelGroup(id string) (model.ChannelGroup, error)
	DeleteChannelGroup(id string) error
	FindAllChannelGroup(app string) ([]model.ChannelGroup, error)
	FindAllChannelGroupByTypefunc(app string, tipo string) ([]model.ChannelGroup, error)
}

