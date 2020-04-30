/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.casousos;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroCasoUso;
import com.gdosoftware.biblioteca.domain.interfaces.ILibroRepository;
import com.gdosoftware.biblioteca.domain.modelo.Libro;
import java.util.List;

/**
 *
 * @author Dani
 * type LibroCasoUsoImpl struct {
 *      repo  ILibroRepository
 * }
 * 
 * (repo *ILibroRepository) func createLibro(libro *model.Libro) (*modelo.Libro, error){
 *      libro:=repo.create(libro)
 *      return libro, nill;
 * }
 * (repo *LibroCasoUsoImpl) func retrieveLibro(id number) (*modelo.Libro, error)
 * (repo *LibroCasoUsoImpl) func updateLibro(id number, libro *model.Libro) (*modelo.Libro, error)
 * (repo *LibroCasoUsoImpl) func deleteLibro(id number) (error)
 * (repo *LibroCasoUsoImpl) func findAllLibro() ([]modelo.Libro, error)
 */
public class LibroCasoUsoImpl implements ILibroCasoUso{
    private ILibroRepository libroRepo;

    public LibroCasoUsoImpl(ILibroRepository libroRepo) {
        this.libroRepo = libroRepo;
    }
    
    

    @Override
    public Libro createLibro(Libro libro) {
        return libroRepo.create(libro);
    }

    @Override
    public Libro retriveLibro(Long id) {
        return libroRepo.retrieve(id);
    }

    @Override
    public Libro updateLibro(Long id, Libro libro) {
        return libroRepo.update(id,libro);
    }

    @Override
    public void deleteLibro(Long id) {
        libroRepo.deltete(id);
    }

    @Override
    public List<Libro> findAllLibro() {
        return libroRepo.findAll();
    }

    
}
