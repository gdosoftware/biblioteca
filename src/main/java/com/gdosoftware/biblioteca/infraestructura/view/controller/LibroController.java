/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.view.controller;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroCasoUso;
import com.gdosoftware.biblioteca.domain.modelo.Libro;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

/**
 *
 * @author Dani
 * type LibroController struct{
 *  libroCasoUso LibroCasoUso
 * }
 * 
 * (cu *LibroController) createLibro(libro *model.Libro): (*model.Libro,error){
 *   cu.createLibro(libro)
 * }
 */
@RestController
public class LibroController {
       
    private ILibroCasoUso libroCasoUso;
    
    @PostMapping("/libro")
    public ResponseEntity<Libro> creteLibro(@RequestBody()Libro libro){
        return ResponseEntity.ok(libroCasoUso.createLibro(libro));
    }
}
