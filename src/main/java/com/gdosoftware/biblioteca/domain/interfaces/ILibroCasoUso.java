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
 * type ILibroCasoUso interface {
 *  createLibro(libro *Libro) (*modelo.Libro,error)
 *  updateLibro(id number, libro *Libro) (*modelo.Libro,error)
 *  retrieveLibro(id number) (modelo.Libro,error)
 *  deleteLibro (id number) error
 *  findAllLibro() ([]modelo.Libro, error)
 * }
 * 
 */
public interface ILibroCasoUso {
    public Libro createLibro(Libro libro);
    public Libro retriveLibro(Long id);
    public Libro updateLibro(Long id, Libro libro);
    public void deleteLibro(Long id);
    public List<Libro> findAllLibro();
}
