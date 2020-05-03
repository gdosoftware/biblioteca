package http

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/usecase"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/infrastructure/helper"
	"github.com/gdosoftware/biblioteca/domain/model"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type ChannelGroupHttp struct {
	logger logger.Logger
	support  *helper.SupportHttp
	caso interfaces.IChannelGroupUseCase
}

func CreateChannelGroupHttp(caso *usecase.ChannelGroupUseCaseImpl) *ChannelGroupHttp {
	return &ChannelGroupHttp{logger: logger.GetDefaultLogger(),
					  support : helper.CreateSupportHttp(false),
		             caso: caso}
}

func (s *ChannelGroupHttp) AltaChannelGroup(w rest.ResponseWriter, r *rest.Request) {
	defer r.Body.Close()

	var toSave model.ChannelGroup
	if err := s.support.ReadBody(&toSave, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}

	insert, err := s.caso.CreateChannelGroup(&toSave)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert ChannelGroup": toSave}).Error("Error saving Channel Group")
		s.support.WriteError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *ChannelGroupHttp) ModificacionChannelGroup(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()

	var toUpdate model.ChannelGroup
	if err := s.support.ReadBody(&toUpdate, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}
	// Udpate a item
	updated, err := s.caso.UpdateChannelGroup(id, &toUpdate)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(updated)
	}
}

func (s *ChannelGroupHttp) BorrarChannelGroup(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()

	err := s.caso.DeleteChannelGroup(id)

	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(id)
	}
}

func (s *ChannelGroupHttp) RecuperarChannelGroup(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	logger.GetDefaultLogger().Infof("Request to get One Channel Group for id", id)

	if id == "" {
		rest.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	s.logger.WithFields(logger.Fields{"id": id}).Debug("Searching for Channel Group with specified Id")

	item, err := s.caso.RetrieveChannelGroup(id)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "id": id}).Error("Getting Channel Group by id")
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(item)
	}
}

func (s *ChannelGroupHttp) RecuperarTodosLosChannelGroups(w rest.ResponseWriter, r *rest.Request) {

	logger.GetDefaultLogger().Infof("Request to get all channel group")

	bins, err := s.caso.FindAllChannelGroup()
	if err != nil {
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(bins)
	}
}
