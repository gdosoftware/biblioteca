/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.interfaces;

import com.gdosoftware.biblioteca.domain.modelo.Prestamo;
import java.util.List;

/**
 *
 * @author Dani
 */
public interface IPrestamoRepository {
    public Prestamo create(Prestamo libro);
    public Prestamo update(Long id, Prestamo libro);
    public Prestamo retrieve (Long id);
    public void deltete (Long id);
    public List<Prestamo> findAll();
}
