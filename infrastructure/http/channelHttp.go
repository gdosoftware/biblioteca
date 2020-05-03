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

const modelName="Channel Group"

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

func (s *ChannelGroupHttp) CreateChannelGroupHttp(w rest.ResponseWriter, r *rest.Request) {
	defer r.Body.Close()

	var toSave model.ChannelGroup
	if err := s.support.ReadBody(&toSave, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}

	insert, err := s.caso.CreateChannelGroup(&toSave)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert ChannelGroup": toSave}).Error("Error saving "+modelName)
		s.support.WriteError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *ChannelGroupHttp) UpdateChannelGroupHttp(w rest.ResponseWriter, r *rest.Request) {
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
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating "+modelName)
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(updated)
	}
}

func (s *ChannelGroupHttp) DeleteChannelGroupHttp(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()

	err := s.caso.DeleteChannelGroup(id)

	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error deleting "+modelName)
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(id)
	}
}

func (s *ChannelGroupHttp) RetrieveChannelGroupHttp(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	logger.GetDefaultLogger().Infof("Request to get One "+modelName+" for id", id)

	if id == "" {
		rest.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	s.logger.WithFields(logger.Fields{"id": id}).Debug("Searching for "+modelName+" with specified Id")

	item, err := s.caso.RetrieveChannelGroup(id)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "id": id}).Error("Getting "+modelName+" by id")
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(item)
	}
}

func (s *ChannelGroupHttp) FindAllChannelGroupHttp(w rest.ResponseWriter, r *rest.Request) {
	app := r.PathParam("app")
	logger.GetDefaultLogger().Infof("Request to get all "+modelName)

	if app == "" {
		rest.Error(w, "App is mandatory", http.StatusBadRequest)
		return
	}

	s.logger.WithFields(logger.Fields{"app": app}).Debug("Searching for " + modelName + " with specified application")

	items, err := s.caso.FindAllChannelGroup(app)
	if err != nil {
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(items)
	}
}

func (s *ChannelGroupHttp) FindAllChannelGroupByTypeHttp(w rest.ResponseWriter, r *rest.Request) {
	app := r.PathParam("app")
	tipo:= r.PathParam("type")
	logger.GetDefaultLogger().Infof("Request to get all " + modelName + " by application and type")

	if app == "" {
		rest.Error(w, "App is mandatory", http.StatusBadRequest)
		return
	}
	if tipo == "" {
		rest.Error(w, "type is mandatory", http.StatusBadRequest)
		return
	}

	s.logger.WithFields(logger.Fields{"app": app, "type":tipo}).Debug("Searching for " + modelName + "  with specified application and type")

	items, err := s.caso.FindAllChannelGroupByType(app, tipo)
	if err != nil {
		s.support.WriteError(err, w)
	} else {
		w.WriteJson(items)
	}
}

