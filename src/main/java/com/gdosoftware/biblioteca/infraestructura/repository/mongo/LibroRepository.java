/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.gdosoftware.biblioteca.infraestructura.repository.mongo;

import com.gdosoftware.biblioteca.domain.interfaces.ILibroRepository;
import com.gdosoftware.biblioteca.domain.modelo.Libro;
import java.util.List;
import org.springframework.jdbc.core.JdbcTemplate;

/**
 *
 package repository

import (
	"errors"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/domain/modelo"
	"github.com/gdosoftware/biblioteca/infraestructura/health"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBLibroRepository struct {
	config         *mgo.DialInfo
	log            logger.Logger
	session        *mgo.Session
	collectionName string
}

func CreateDBLibroRepository(config *mgo.DialInfo, log logger.Logger, collectionName string) (*DBLibroRepository, error) {
	if collectionName == "" {
		return nil, errors.New("collection name is mandatory")
	}
	log.Infof("Connecting to database with config %v", config)
	session, err := mgo.DialWithInfo(config)
	if err != nil {
		return nil, err
	}
	return &DBLibroRepository{config: config, log: log, session: session, collectionName: collectionName}, nil
}

func (d *DBLibroRepository) Create(toSave *modelo.Libro) (*modelo.Libro, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()

	toSave.Id = bson.NewObjectId()

	d.log.WithFields(logger.Fields{"Id": toSave.Id.Hex()}).Debug("Saving/Updating Libro")
	_, err := sessionCopy.DB(d.config.Database).C(d.collectionName).UpsertId(toSave.Id, toSave)
	if err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error occur while trying to Insert a bin")
		return toSave, err
	}

	return toSave, nil

}

func (d *DBLibroRepository) Update(id string, toUpdate *modelo.Libro) (*modelo.Libro, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()

	toUpdate.Id = bson.ObjectIdHex(id)

	d.log.WithFields(logger.Fields{"Id": toUpdate.Id.Hex()}).Debug("Saving/Updating bin")
	_, err := sessionCopy.DB(d.config.Database).C(d.collectionName).UpsertId(toUpdate.Id, toUpdate)
	if err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error occur while trying to Insert a bin")
		return toUpdate, err
	}

	return toUpdate, nil
}

func (d *DBLibroRepository) Delete(id string) error {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()

	toDelete, _ := d.Retrieve(id)

	toDelete.Deleted = "true"

	d.log.WithFields(logger.Fields{"Id": toDelete.Id.Hex()}).Debug("deleting bin")
	_, err := sessionCopy.DB(d.config.Database).C(d.collectionName).UpsertId(toDelete.Id, toDelete)
	if err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error occur while trying to Insert a bin")
		return err
	}

	return nil
}

func (d *DBLibroRepository) Retrieve(id string) (modelo.Libro, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbLibro modelo.Libro
	if err := d.createQuery(sessionCopy, bson.M{"_id": bson.ObjectIdHex(id)}).One(&dbLibro); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return dbLibro, errors.New("Error in query execution")
	}
	return dbLibro, nil
}


func (d *DBLibroRepository) FindAll() ([]modelo.Libro, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbLibro []modelo.Libro
	if err := d.createQuery(sessionCopy, bson.M{"deleted": "false"}).All(&dbLibro); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return nil, errors.New("Error in query execution")
	}

	return dbLibro, nil
}

func (d *DBLibroRepository) createQuery(session *mgo.Session, criteria bson.M) *mgo.Query {
	return session.DB(d.config.Database).C(d.collectionName).Find(criteria)
}

func (d *DBLibroRepository) Name() string {
	return "MongoDBLibroRepository"
}

func (d *DBLibroRepository) Health() error {
	return d.session.Ping()
}

func (d *DBLibroRepository) Stats() interface{} {
	buildInfo, _ := d.session.BuildInfo()
	return health.RepoStats{
		BuildInfo:   buildInfo,
		LiveServers: d.session.LiveServers(),
	}
}

 * 
 */
public class LibroRepository implements ILibroRepository{
    
    private JdbcTemplate template;

    public LibroRepository(JdbcTemplate template) {
        this.template = template;
    }
    
    

    @Override
    public Libro create(Libro libro) {
        //template.
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Libro update(Long id, Libro libro) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Libro retrieve(Long id) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public void deltete(Long id) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public List<Libro> findAll() {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }
    
}
