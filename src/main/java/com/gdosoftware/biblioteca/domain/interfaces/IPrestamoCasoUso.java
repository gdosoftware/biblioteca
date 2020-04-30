/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.interfaces;

import com.gdosoftware.biblioteca.domain.modelo.Prestamo;

/**
 *
 * @author Dani
 */
public interface IPrestamoCasoUso {
    public Prestamo createPrestamo(Long socioId, Long libroId);
    public Prestamo devolucionPrestamo(Long prestamoId);
}
