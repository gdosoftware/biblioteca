package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/casousos"
	"github.com/gdosoftware/biblioteca/domain/modelo"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type SocioApi struct {
	logger  logger.Logger
	support *SupportAPI
	caso *casousos.SocioCasoUsoImpl
}


func createSocioApi(caso *casousos.SocioCasoUsoImpl) *SocioApi{
    retrun &LSocioApi{logger :  logger.GetDefaultLogger(), 
					support : api.CreateSupportAPI(),
					caso: caso}
}

func (s *SocioApi) altaSocio(w rest.ResponseWriter, r *rest.Request) {
	defer r.Body.Close()

	var toSave modelo.Socio
	if err := s.support.readBody(&toSave, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}

	insert, err := s.caso.CreateSocio(&toSave)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": libro}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *SocioApi) modificacionSocio(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()
	
	var toUpdate modelo.Socio
	if err := s.support.readBody(&toUpdate, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}
	// Udpate a item
	updated, err := s.caso.updateSocio(id, &toUpdate)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.support.writeError(err, w)
	} else {
		w.WriteJson(updated)
	}
}

func (s *SocioApi) borrarSocio(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()
	
	err := s.caso.DeleteSocio(id)

	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.support.writeError(err, w)
	} else {
		w.WriteJson(id)
	}
}

func (s *SocioApi) recuperarSocio(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	logger.GetDefaultLogger().Infof("Request to get One Channel Group for id", id)

	if id == "" {
		rest.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}
	
	s.logger.WithFields(logger.Fields{"id": id}).Debug("Searching for Channel Group with specified Id")

	item, err := s.caso.findSocioById(id)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "id": id}).Error("Getting Channel Group by id")
		s.support.writeError(err, w)
	} else {
		w.WriteJson(item)
	}
}

func (s *SocioApi) recuperarTodosLosSocios(w rest.ResponseWriter, r *rest.Request) {
	
	logger.GetDefaultLogger().Infof("Request to get all channel group")

	
	socios, err := s.caso.findAllSocio()
	if err != nil {
		s.support.writeError(err, w)
	} else {
		w.WriteJson(socios)
	}
}




