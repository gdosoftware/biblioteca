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
 package casousos

import (
	"github.com/gdosoftware/biblioteca/domain/interfaces"
	"github.com/gdosoftware/biblioteca/domain/modelo"
)

type LibroCasoUsoImpl struct {
	Repo interfaces.ILibroRepository
}

func (r *LibroCasoUsoImpl) CreateLibro(libro *modelo.Libro) (*modelo.Libro, error) {
	return r.Repo.Create(libro)
}

func (r *LibroCasoUsoImpl) RetrieveLibro(id string) (modelo.Libro, error) {
	return r.Repo.Retrieve(id)
}

func (r *LibroCasoUsoImpl) UpdateLibro(id string, libro *modelo.Libro) (*modelo.Libro, error) {
	return r.Repo.Update(id, libro)
}

func (r *LibroCasoUsoImpl) DeleteLibro(id string) error {
	return r.Repo.Delete(id)
}

func (r *LibroCasoUsoImpl) FindAllLibro() ([]modelo.Libro, error) {
	return r.Repo.FindAll()
}

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
