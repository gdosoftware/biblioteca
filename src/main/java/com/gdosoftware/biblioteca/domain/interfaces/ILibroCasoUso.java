/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.interfaces;

import com.gdosoftware.biblioteca.domain.modelo.Libro;
import java.util.List;

/**
 *
 package interfaces

import "github.com/gdosoftware/biblioteca/domain/modelo"

type ILibroCasoUso interface {
	CreateLibro(libro *modelo.Libro) (*modelo.Libro, error)
	UpdateLibro(id string, libro *modelo.Libro) (*modelo.Libro, error)
	RetrieveLibro(id string) (modelo.Libro, error)
	DeleteLibro(id string) error
	FindAllLibro() ([]modelo.Libro, error)
}
 * 
 */
public interface ILibroCasoUso {
    public Libro createLibro(Libro libro);
    public Libro retriveLibro(Long id);
    public Libro updateLibro(Long id, Libro libro);
    public void deleteLibro(Long id);
    public List<Libro> findAllLibro();
}
