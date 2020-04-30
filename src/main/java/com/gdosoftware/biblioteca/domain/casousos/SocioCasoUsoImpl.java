/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.casousos;

import com.gdosoftware.biblioteca.domain.interfaces.ISocioCasoUso;
import com.gdosoftware.biblioteca.domain.interfaces.ISocioRepository;
import com.gdosoftware.biblioteca.domain.modelo.Socio;
import java.util.List;

/**
 *
 * @author Dani
 */
public class SocioCasoUsoImpl implements ISocioCasoUso{
    
    private ISocioRepository socioRepo;

    @Override
    public Socio createLibro(Socio socio) {
        return socioRepo.create(socio);
    }

    @Override
    public Socio retriveLibro(Long id) {
        return socioRepo.retrieve(id);
    }

    @Override
    public Socio updateLibro(Long id, Socio socio) {
        return socioRepo.update(id, socio);
    }

    @Override
    public void deleteLibro(Long id) {
        socioRepo.deltete(id);
    }

    @Override
    public List<Socio> findAllLibro() {
        return socioRepo.findAll();
    }
    
}
