package http

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/casousos"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type LibroHttp struct {
	logger logger.Logger
	SupportAPI
	caso interfaces.ILibroCasoUso
}

func CreateLibroHttp(caso *casousos.LibroCasoUsoImpl) *LibroHttp {
	return &LibroHttp{logger: logger.GetDefaultLogger(),
		caso: caso}
}

func (s *LibroHttp) AltaLibro(w rest.ResponseWriter, r *rest.Request) {
	defer r.Body.Close()

	var toSave modelo.Libro
	if err := s.readBody(&toSave, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}

	insert, err := s.caso.CreateLibro(&toSave)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": toSave}).Error("Error saving Channel Group")
		s.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *LibroHttp) ModificacionLibro(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()

	var toUpdate modelo.Libro
	if err := s.readBody(&toUpdate, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}
	// Udpate a item
	updated, err := s.caso.UpdateLibro(id, &toUpdate)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.writeError(err, w)
	} else {
		w.WriteJson(updated)
	}
}

func (s *LibroHttp) BorrarLibro(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	defer r.Body.Close()

	err := s.caso.DeleteLibro(id)

	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err}).Error("Error updating item")
		s.writeError(err, w)
	} else {
		w.WriteJson(id)
	}
}

func (s *LibroHttp) RecuperarLibro(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	logger.GetDefaultLogger().Infof("Request to get One Channel Group for id", id)

	if id == "" {
		rest.Error(w, "Id is mandatory", http.StatusBadRequest)
		return
	}

	s.logger.WithFields(logger.Fields{"id": id}).Debug("Searching for Channel Group with specified Id")

	item, err := s.caso.RetrieveLibro(id)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "id": id}).Error("Getting Channel Group by id")
		s.writeError(err, w)
	} else {
		w.WriteJson(item)
	}
}

func (s *LibroHttp) RecuperarTodosLosLibros(w rest.ResponseWriter, r *rest.Request) {

	logger.GetDefaultLogger().Infof("Request to get all channel group")

	bins, err := s.caso.FindAllLibro()
	if err != nil {
		s.writeError(err, w)
	} else {
		w.WriteJson(bins)
	}
}
