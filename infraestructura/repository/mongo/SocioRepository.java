/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.repository.mongo;

import com.gdosoftware.biblioteca.domain.interfaces.ISocioRepository;
import com.gdosoftware.biblioteca.domain.modelo.Socio;
import java.util.List;

/**
 *
 * @author Dani
 */
public class SocioRepository implements ISocioRepository{

    @Override
    public Socio create(Socio libro) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Socio update(Long id, Socio libro) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Socio retrieve(Long id) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public void deltete(Long id) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public List<Socio> findAll() {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }
    
}
