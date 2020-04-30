package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type LibroApi struct {
	logger  logger.Logger
	support *SupportAPI
	caso *interfaces.ILibroCasoUso
}


func createLibroApi(logger  logger.Logger){
    retrun &LibroApi(logger :  logger.GetDefaultLogger())
}

func (s *LibroApi) AltaLibro(w rest.ResponseWriter, r *rest.Request) {
	defer s.Body.Close()

	var toSAve modelo.Libro
	if err := s.support.readBody(&toSave, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}

	insert, err := s.caso.CreateLibro(&libro)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": libro}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *LibroApi) ModificacionLibro(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()
	
	var toUpdate modelo.Libro
	if err := s.support.readBody(&toUpdate, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}
	// Udpate a item
	updated, err := s.caso.updateLibro(id, &toUpdate)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.support.writeError(err, w)
	} else {
		w.WriteJson(updated)
	}
}

func (s *LibroApi) BorrarLibro(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()
	
	err := s.caso.DeleteLibro(id)

	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.support.writeError(err, w)
	} else {
		w.WriteJson(id)
	}
}

func (s *LibroApi) RecuperarLibro(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	logger.GetDefaultLogger().Infof("Request to get One Channel Group for id", id)

	if id == "" {
		rest.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}
	
	s.logger.WithFields(logger.Fields{"id": id}).Debug("Searching for Channel Group with specified Id")

	item, err := s.caso.findLibroById(id)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "id": id}).Error("Getting Channel Group by id")
		s.support.writeError(err, w)
	} else {
		w.WriteJson(item)
	}
}

func (s *LibroApi) RecuperarTodosLosLibros() ((w rest.ResponseWriter, r *rest.Request) {
	
	logger.GetDefaultLogger().Infof("Request to get all channel group")

	
	bins, err := s.caso.findAllLibro()
	if err != nil {
		s.support.writeError(err, w)
	} else {
		w.WriteJson(bins)
	}
}




