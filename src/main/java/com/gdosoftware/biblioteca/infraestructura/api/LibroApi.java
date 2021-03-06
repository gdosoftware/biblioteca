/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.api;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroCasoUso;
import com.gdosoftware.biblioteca.domain.modelo.Libro;
import java.util.List;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

/**
 *
 package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gdosoftware/biblioteca/domain/casousos"
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

type LibroApi struct {
	logger logger.Logger
	SupportAPI
	caso interfaces.ILibroCasoUso
}

func CreateLibroApi(caso *casousos.LibroCasoUsoImpl) *LibroApi {
	return &LibroApi{logger: logger.GetDefaultLogger(),
		caso: caso}
}

func (s *LibroApi) AltaLibro(w rest.ResponseWriter, r *rest.Request) {
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

func (s *LibroApi) ModificacionLibro(w rest.ResponseWriter, r *rest.Request) {
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

func (s *LibroApi) BorrarLibro(w rest.ResponseWriter, r *rest.Request) {
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

func (s *LibroApi) RecuperarLibro(w rest.ResponseWriter, r *rest.Request) {
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

func (s *LibroApi) RecuperarTodosLosLibros(w rest.ResponseWriter, r *rest.Request) {

	logger.GetDefaultLogger().Infof("Request to get all channel group")

	bins, err := s.caso.FindAllLibro()
	if err != nil {
		s.writeError(err, w)
	} else {
		w.WriteJson(bins)
	}
}

 * }
 */
public class LibroApi {
       
    private ILibroCasoUso libroCasoUso;
    
    public LibroApi(ILibroCasoUso libroCasoUso) {
        this.libroCasoUso = libroCasoUso;
    }
        
    public Libro altaLibro(Libro libro){
        return libroCasoUso.createLibro(libro);
    }
    
    public Libro modificacionLibro(Long id, Libro libro){
        return libroCasoUso.updateLibro(id, libro);
    }
    
    public void borrarLibro(Long id){
        libroCasoUso.deleteLibro(id);
    }
    
    public Libro recuperarLibro(Long id){
        return libroCasoUso.retriveLibro(id);
    }
    
    public List<Libro> recuperarTodosLosLibro(){
        return libroCasoUso.findAllLibro();
    }
}
