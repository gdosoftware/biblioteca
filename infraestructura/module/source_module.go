package module

import (
	"fmt"
	"os"
	"strings"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/infraestructura/repository/mongo"
	"github.com/gdosoftware/biblioteca/infraestructura/helper"

	"gopkg.in/mgo.v2"
)

const (
	DefaultSkuCollectionName = "sku"

	// Environment
	MongoDBBLibroCollectionName = "MONGO_LIBRO_COLLECTION"
	MongoDBSocioCollectionName = "MONGO_SOCIO_COLLECTION"
	MongoDBPrestamoCollectionName = "MONGO_PRESTAMO_COLLECTION"
	
	MongoDBName                           = "MONGO_DATABASE_NAME"
	MongoDBAddress                        = "MONGO_DATABASE_ADDR"
	JWTPublicKeyEnv                       = "JWT_PUBLIC_KEY"
	JWTTokenTask                          = "JWT_TOKEN_TASK"
)

type SourceFactory struct {
}

var mandatoryEnvironmentVariables = []string{
	MongoDBName,
	MongoDBAddress,
}

const DefaultLibroCollectionName = "libro"
const DefaultSocioCollectionName = "socio"
const DefaultPrestamoCollectionName = "prestamo"


func NewSourceFactory() error {
	if missingVariables := validateEnvironmentVariables(); len(missingVariables) > 0 {
		return fmt.Errorf("The following environment variables are missing and the application can't start: %v", missingVariables)
	}
	return nil
}

func createMongoConfig(log logger.Logger) *mgo.DialInfo {
	databaseName, dataBaseAddr := os.Getenv("MONGO_DATABASE_NAME"), os.Getenv("MONGO_DATABASE_ADDR")
	if databaseName == "" || dataBaseAddr == "" {
		log.Fatal("Database name/address is mandatory")
		return nil
	}
	return &mgo.DialInfo{
		Addrs:    strings.Split(dataBaseAddr, ","),
		Database: databaseName,
	}
}

func getEnvOrDefault(envName string, defaultValue string) string {
	if env := os.Getenv(envName); env == "" {
		return defaultValue
	} else {
		return env
	}
}

func CreateLibroRepository(log logger.Logger) *repository.DBLibroRepository {
	collectionName := getEnvOrDefault("MONGO_LIBRO_COLLECTION", DefaultLibroCollectionName)
	repo, err := repository.CreateDBLibroRepository(createMongoConfig(log), log, collectionName)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating bin repository")
		return nil
	}
	return repo
}
/*
func CreateSocioRepository(log logger.Logger) *repository.DBSocioRepository {
	collectionName := getEnvOrDefault("MONGO_SOCIO_COLLECTION", DefaultSocioCollectionName)
	repo, err := repository.CreateDBSocioRepository(createMongoConfig(log), log, collectionName)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating bin repository")
		return nil
	}
	return repo
}

func CreatePrestamoRepository(log logger.Logger) *repository.DBPrestamoRepository {
	collectionName := getEnvOrDefault("MONGO_PRESTAMO_COLLECTION", DefaultPrestamoCollectionName)
	repo, err := repository.CreateDBPrestamoRepository(createMongoConfig(log), log, collectionName)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating bin repository")
		return nil
	}
	return repo
}*/


func CreateJwtDecoder() *helper.JwtDecoder {
	return helper.CreateJwtDecoder(os.Getenv(JWTPublicKeyEnv))
}

func GetTokenTask() string {
	return os.Getenv(JWTTokenTask)
}



func validateEnvironmentVariables() []string {
	var missingEnvVars []string
	for _, v := range mandatoryEnvironmentVariables {
		if strings.TrimSpace(os.Getenv(v)) == "" {
			missingEnvVars = append(missingEnvVars, v)
		}
	}
	return missingEnvVars
}

