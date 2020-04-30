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
public interface ISocioCasoUso {
    public Socio createLibro(Socio libro);
    public Socio retriveLibro(Long id);
    public Socio updateLibro(Long id, Socio libro);
    public void deleteLibro(Long id);
    public List<Socio> findAllLibro();
}
