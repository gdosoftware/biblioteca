package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/casousos"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type PrestamoApi struct {
	logger  logger.Logger
	support *SupportAPI
	caso    *casousos.PrestamoCasoUsoImpl
}

func createPrestamoApi(caso *casousos.PrestamoCasoUsoImpl) *PrestamoApi {
	return &PrestamoApi{logger: logger.GetDefaultLogger(),
		support: CreateSupportAPI(),
		caso:    caso}
}

func (s *PrestamoApi) prestarLibro(w rest.ResponseWriter, r *rest.Request) {
	idLibro := r.PathParam("libroId")
	idSocio := r.PathParam("socioId")
	defer r.Body.Close()

	insert, err := s.caso.CreatePrestamo(idLibro, idSocio)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": insert}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *PrestamoApi) devolverLibro(w rest.ResponseWriter, r *rest.Request) {
	idPrestamo := r.PathParam("prestamoId")
	defer r.Body.Close()

	insert, err := s.caso.UpdatePrestamo(idPrestamo)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": insert}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}
