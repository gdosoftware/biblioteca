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


func createLibroApi(support *SupportAPI, casoLibros interfaces.ILibroCasoUso){
    retrun &LibroApi(logger :  logger.GetDefaultLogger())
}

func (s *LibroApi) altaLIbro(w rest.ResponseWriter, r *rest.Request) {
	defer s.Body.Close()

	var libro modelo.Libro
	if err := s.support.readBody(&toSave, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(err.Error())
		return
	}

	insert, err := s.CreateLibro(&libro)
	if err != nil {
		s.logger.WithFields(logger.Fields{"error": err, "insert Libro": libro}).Error("Error saving Channel Group")
		s.support.writeError(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.WriteJson(insert)
	}
}

func (l *LibroApi) altaLibro(libro *modelo.Libro) (*modelo.Libro, error) {
	return l.caso.CreateLibro(libro)
}

func (l *LibroApi) modificacionLibro(id string, libro *modelo.Libro) (*modelo.Libro, error) {
	return l.caso.UpdateLibro(id, libro)
}

func (l *LibroApi) borrarLibro(id string) error {
	return l.caso.DeleteLibro(id)
}

func (l *LibroApi) recuperarLibro(id string) (modelo.Libro, error) {
	return l.caso.RetrieveLibro(id)
}

func (l *LibroApi) recuperarTodosLosLibros() ([]modelo.Libro, error) {
	return l.caso.FindAllLibro()
}
