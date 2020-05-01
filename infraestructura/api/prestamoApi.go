package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type PrestanoApi struct {
	logger  logger.Logger
	support *SupportAPI
	caso *interfaces.IPrestamoCasoUso
}


func createPrestamoApi(logger  logger.Logger, caso *interfaces.IPrestamoCasoUso){
    retrun &PrestamoApi{logger :  logger.GetDefaultLogger(), 
					support : api.CreateSupportAPI(),
					caso: caso}
}

func (s *PrestanoApi) prestarLibro(w rest.ResponseWriter, r *rest.Request) {
	idLibro := r.PathParam("libroId")
	idSocio := r.PathParam("socioId")
	defer s.Body.Close()


	insert, err := s.caso.CreatePrestamo(idLibro, idSocio)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": libro}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (s *PrestanoApi) devolverLibro(w rest.ResponseWriter, r *rest.Request) {
	idPrestamo := r.PathParam("prestamoId")
	defer s.Body.Close()

	insert, err := s.caso.UpdatePrestamo(idPrestamo)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": libro}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

