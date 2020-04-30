/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.api;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroCasoUso;
import com.gdosoftware.biblioteca.domain.modelo.Libro;

/**
 *
 * @author Dani
 * type LibroApi struct{
 *  libroCasoUso ILibroCasoUso
 * }
 * 
 * (cu *LibroApi) createLibro(libro *model.Libro): (*model.Libro,error){
 *   cu.createLibro(libro)
 * }
 */
public class LibroApi {
       
    private ILibroCasoUso libroCasoUso;

    public LibroApi(ILibroCasoUso libroCasoUso) {
        this.libroCasoUso = libroCasoUso;
    }
        
    public Libro crete(Libro libro){
        return libroCasoUso.createLibro(libro);
    }
}
