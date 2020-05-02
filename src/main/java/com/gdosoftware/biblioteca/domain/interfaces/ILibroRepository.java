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

type ILibroRepository interface {
	Create(libro *modelo.Libro) (*modelo.Libro, error)
	Update(id string, libro *modelo.Libro) (*modelo.Libro, error)
	Retrieve(id string) (modelo.Libro, error)
	Delete(id string) error
	FindAll() ([]modelo.Libro, error)
}
 */
public interface ILibroRepository {
    public Libro create(Libro libro);
    public Libro update(Long id, Libro libro);
    public Libro retrieve (Long id);
    public void deltete (Long id);
    public List<Libro> findAll();
}
