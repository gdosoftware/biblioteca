package repository

import (
	"errors"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/domain/model"
	"github.com/gdosoftware/biblioteca/infrastructure/health"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBChannelGroupRepository struct {
	config         *mgo.DialInfo
	log            logger.Logger
	session        *mgo.Session
	collectionName string
}

func CreateDBChannelGroupRepository(config *mgo.DialInfo, log logger.Logger, collectionName string) (*DBChannelGroupRepository, error) {
	if collectionName == "" {
		return nil, errors.New("collection name is mandatory")
	}
	log.Infof("Connecting to database with config %v", config)
	session, err := mgo.DialWithInfo(config)
	if err != nil {
		return nil, err
	}
	return &DBChannelGroupRepository{config: config, log: log, session: session, collectionName: collectionName}, nil
}

func (d *DBChannelGroupRepository) Create(toSave *model.ChannelGroup) (*model.ChannelGroup, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()

	toSave.Id = bson.NewObjectId()

	d.log.WithFields(logger.Fields{"Id": toSave.Id.Hex()}).Debug("Saving/Updating ChannelGroup")
	_, err := sessionCopy.DB(d.config.Database).C(d.collectionName).UpsertId(toSave.Id, toSave)
	if err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error occur while trying to Insert a bin")
		return toSave, err
	}

	return toSave, nil

}

func (d *DBChannelGroupRepository) Update(id string, toUpdate *model.ChannelGroup) (*model.ChannelGroup, error) {
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

func (d *DBChannelGroupRepository) Delete(id string) error {
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

func (d *DBChannelGroupRepository) Retrieve(id string) (model.ChannelGroup, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbChannelGroup model.ChannelGroup
	if err := d.createQuery(sessionCopy, bson.M{"_id": bson.ObjectIdHex(id)}).One(&dbChannelGroup); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return dbChannelGroup, errors.New("Error in query execution")
	}
	return dbChannelGroup, nil
}


func (d *DBChannelGroupRepository) FindAll(app string) ([]model.ChannelGroup, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbChannelGroup []model.ChannelGroup
	if err := d.createQuery(sessionCopy, bson.M{"deleted": "false","application":app}).All(&dbChannelGroup); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return nil, errors.New("Error in query execution")
	}

	return dbChannelGroup, nil
}


func (d *DBChannelGroupRepository) FindByType(app string, tipo string) ([]model.ChannelGroup, error) {
	sessionCopy := d.session.Copy()
	defer sessionCopy.Close()
	var dbChannelGroup []model.ChannelGroup
	if err := d.createQuery(sessionCopy, bson.M{"deleted": "false","application":app,"type":tipo}).All(&dbChannelGroup); err != nil {
		d.log.WithFields(logger.Fields{"error": err}).Error("Error in query execution")
		return nil, errors.New("Error in query execution")
	}

	return dbChannelGroup, nil
}

func (d *DBChannelGroupRepository) createQuery(session *mgo.Session, criteria bson.M) *mgo.Query {
	return session.DB(d.config.Database).C(d.collectionName).Find(criteria)
}

func (d *DBChannelGroupRepository) Name() string {
	return "MongoDBChannelGroupRepository"
}

func (d *DBChannelGroupRepository) Health() error {
	return d.session.Ping()
}

func (d *DBChannelGroupRepository) Stats() interface{} {
	buildInfo, _ := d.session.BuildInfo()
	return health.RepoStats{
		BuildInfo:   buildInfo,
		LiveServers: d.session.LiveServers(),
	}
}
