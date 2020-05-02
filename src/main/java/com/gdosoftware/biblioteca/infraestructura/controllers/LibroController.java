/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.controllers;

import com.gdosoftware.biblioteca.domain.modelo.Libro;
import com.gdosoftware.biblioteca.infraestructura.api.LibroApi;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

/**
 *
 * @author Dani
 */
@RestController
public class LibroController {
    
    @Autowired
    private LibroApi libroApi;
    
    @PostMapping("/libro")
    public ResponseEntity<Libro> post(@RequestBody()Libro libro){
        return ResponseEntity.ok(libroApi.altaLibro(libro));
    }
    
     @PutMapping("/libro")
    public ResponseEntity<Libro> put(@PathVariable("id")Long id, 
                                     @RequestBody()Libro libro){
        return ResponseEntity.ok(libroApi.modificacionLibro(id, libro));
    }
    
     @DeleteMapping("/libro")
    public ResponseEntity<Libro> delete(@PathVariable("id")Long id){
        libroApi.borrarLibro(id);
        return ResponseEntity.ok().build();
    }
    
     @PutMapping("/libro")
    public ResponseEntity<Libro> getOne(@PathVariable("id")Long id){
        return ResponseEntity.ok(libroApi.recuperarLibro(id));
    }
    
    @PutMapping("/libro")
    public ResponseEntity<List<Libro>> getAll(){
        return ResponseEntity.ok(libroApi.recuperarTodosLosLibro());
    }
}
