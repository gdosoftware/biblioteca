package usecase

import (
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/domain/interfaces"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/domain/model"
)

type ChannelGroupUseCaseImpl struct {
	Repo interfaces.IChannelGroupRepository
}

func CreateChannelGroupUseCaseImpl(repo interfaces.IChannelGroupRepository) *ChannelGroupUseCaseImpl{
	return &ChannelGroupUseCaseImpl{Repo:repo}
}

func (r *ChannelGroupUseCaseImpl) CreateChannelGroup(channelGroup *model.ChannelGroup) (*model.ChannelGroup, error) {
	return r.Repo.Create(channelGroup)
}

func (r *ChannelGroupUseCaseImpl) RetrieveChannelGroup(id string) (model.ChannelGroup, error) {
	return r.Repo.Retrieve(id)
}

func (r *ChannelGroupUseCaseImpl) UpdateChannelGroup(id string, channelGroup *model.ChannelGroup) (*model.ChannelGroup, error) {
	return r.Repo.Update(id, channelGroup)
}

func (r *ChannelGroupUseCaseImpl) DeleteChannelGroup(id string) error {
	return r.Repo.Delete(id)
}

func (r *ChannelGroupUseCaseImpl) FindAllChannelGroup(app string) ([]model.ChannelGroup, error){
	return r.Repo.FindAll(app)
}

func (r *ChannelGroupUseCaseImpl) FindAllChannelGroupByType(app string, tipo string) ([]model.ChannelGroup, error){
	return r.Repo.FindByType(app, tipo)
}

/*
var channelGroup ChannelGroupUseCaseImpl

channelGroup = ChannelGroupUseCaseImpl {Repo : ChannelGroupRepositoryImpl}
channelGroup.CreateLIbro
	 .RetrieveChannelGroup
	 .UpdateChannelGroup
	 .DelteChannelGroup
	 .FindAllChannelGroup

func(l IchannelGroupUseCase)

*/
