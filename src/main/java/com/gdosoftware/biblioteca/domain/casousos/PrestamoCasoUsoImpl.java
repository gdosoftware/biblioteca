/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.casousos;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroRepository;
import com.gdosoftware.biblioteca.domain.interfaces.IPrestamoCasoUso;
import com.gdosoftware.biblioteca.domain.interfaces.IPrestamoRepository;
import com.gdosoftware.biblioteca.domain.interfaces.ISocioRepository;
import com.gdosoftware.biblioteca.domain.modelo.Libro;
import com.gdosoftware.biblioteca.domain.modelo.Prestamo;
import com.gdosoftware.biblioteca.domain.modelo.Socio;

/**
 *
 * @author Dani
 */
public class PrestamoCasoUsoImpl implements IPrestamoCasoUso{
    private ISocioRepository socioRepo;
    private ILibroRepository libroRepo;
    private IPrestamoRepository prestamoRepo;

    @Override
    public Prestamo createPrestamo(Long socioId, Long libroId) {
        Socio socio = socioRepo.retrieve(socioId);
        Libro libro = libroRepo.retrieve(libroId);
        Prestamo prestamo = new Prestamo();
        //add socio y libro
        return prestamoRepo.create(prestamo);
    }

    @Override
    public Prestamo devolucionPrestamo(Long prestamoId) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }
}
