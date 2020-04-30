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
 * @author Dani
 * 
 * type LibroRepository interface{
 * 
 *  create(libro *Libro) (*modelo.Libro,error)
 *  update(id number, libro *Libro) (*modelo.Libro,error)
 *  retrieve(id number) (modelo.Libro,error)
 *  delete (id number) error
 *  findAll() ([]modelo.Libro, error)
 * }
 */
public interface ILibroRepository {
    public Libro create(Libro libro);
    public Libro update(Long id, Libro libro);
    public Libro retrieve (Long id);
    public void deltete (Long id);
    public List<Libro> findAll();
}
