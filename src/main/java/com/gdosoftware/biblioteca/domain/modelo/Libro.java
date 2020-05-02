/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.domain.modelo;

/**
 *
 * package modelo

import "gopkg.in/mgo.v2/bson"

type Libro struct {
	Id        bson.ObjectId `bson:"_id"`
	Titulo    string        `bson:"titulo"`
	Autor     string        `bson:"autor"`
	Stock     int           `bson:"stock"`
	Borrowed int            `bson:"borrowed"`
	Deleted   string        `bson:"deleted"`
}

 * 
 */
public class Libro {
    
    private Long id;
    private String title;
    private String author;
    private int stock;
    private int borrowed;
    private String delete;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getAuthor() {
        return author;
    }

    public void setAuthor(String author) {
        this.author = author;
    }

    public int getStock() {
        return stock;
    }

    public void setStock(int stock) {
        this.stock = stock;
    }

    public int getBorrowed() {
        return borrowed;
    }

    public void setBorrowed(int borrowed) {
        this.borrowed = borrowed;
    }

    public String getDelete() {
        return delete;
    }

    public void setDelete(String delete) {
        this.delete = delete;
    }
    
    
    
}
