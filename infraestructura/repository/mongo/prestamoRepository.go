package repository

import (
	"errors"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/domain/modelo"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBPrestamoRepository struct {
	config         *mgo.DialInfo
	log            logger.Logger
	session        *mgo.Session
	collectionName string
}

func CreateDBPrestamoRepository(config *mgo.DialInfo, log logger.Logger, collectionName string) (*DBPrestamoRepository, error) {
	if collectionName == "" {
		return nil, errors.New("collection name is mandatory")
	}
	log.Infof("Connecting to database with config %v", config)
	session, err := mgo.DialWithInfo(config)
	if err != nil {
		return nil, err
	}
	return &DBPrestamoRepository{config: config, log: log, session: session, collectionName: collectionName}, nil
}
/*
    Create(libro *modelo.Libro) (*modelo.Libro, error)
	Update(id string, libro *modelo.Libro) (*modelo.Libro, error)
	Retrieve(id string) (modelo.Libro, error)
	Delete(id string) error
	FindAll() ([]modelo.Libro, error)
*/


func (d *DBPrestamoRepository) Create(toSave *modelo.Prestamo) (*modelo.Prestamo, error) {
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

func (d *DBPrestamoRepository) Update(id string, toUpdate *modelo.Prestamo) (*model.Prestamo, error) {
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

func (d *DBPrestamoRepository) Delete(id string) error {
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

func (d *DBPrestamoRepository) Retrieve(id string) (modelo.Prestamo, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbPrestamoo modelo.Prestamo
	if err := d.createQuery(sessionCopy, bson.M{"_id": bson.ObjectIdHex(id)}).One(&dbPrestamoo); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return dbPrestamoo, errors.New("Error in query execution")
	}
	return dbLibro, nil
}


func (d *DBPrestamoRepository) FindAll() ([]modelo.Prestamo, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbPrestamo []modelo.Prestamo
	if err := d.createQuery(sessionCopy, bson.M{"deleted": "false"}).All(&dbPrestamo); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return nil, errors.New("Error in query execution")
	}

	return dbPrestamo, nil
}

func (d *DBPrestamoRepository) createQuery(session *mgo.Session, criteria bson.M) *mgo.Query {
	return session.DB(d.config.Database).C(d.collectionName).Find(criteria)
}

func (d *DBPrestamoRepository) Name() string {
	return "MongoDBPrestamoRepository"
}

func (d *DBPrestamoRepository) Health() error {
	return d.session.Ping()
}

func (d *DBPrestamoRepository) Stats() interface{} {
	buildInfo, _ := d.session.BuildInfo()
	return health.RepoStats{
		BuildInfo:   buildInfo,
		LiveServers: d.session.LiveServers(),
	}
}
