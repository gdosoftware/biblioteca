/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.repository.postgres;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroRepository;
import com.gdosoftware.biblioteca.domain.modelo.Libro;
import java.util.List;

/**
 *
 * @author Dani
 */
public class LibroRepository implements ILibroRepository
{

    @Override
    public Libro create(Libro libro) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Libro update(Long id, Libro libro) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Libro retrieve(Long id) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public void deltete(Long id) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public List<Libro> findAll() {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }
    
}
