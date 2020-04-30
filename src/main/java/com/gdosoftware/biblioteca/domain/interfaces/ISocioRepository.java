/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.interfaces;

import com.gdosoftware.biblioteca.domain.modelo.Socio;
import java.util.List;

/**
 *
 * @author Dani
 */
public interface ISocioRepository {
    public Socio create(Socio libro);
    public Socio update(Long id, Socio libro);
    public Socio retrieve (Long id);
    public void deltete (Long id);
    public List<Socio> findAll();
}
