package module

import (
	"fmt"
	"os"
	"strings"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/infrastructure/repository/mongo"
	"gitlab.com/fravega-it/adn/ipos/configuracion/agrupaciones/infrastructure/helper"

	"gopkg.in/mgo.v2"
)

const (
	DefaultSkuCollectionName = "sku"

	// Environment
	MongoDBBChannelGroupCollectionName = "MONGO_LIBRO_COLLECTION"
	
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

const DefaultChannelGroupCollectionName = "channelGroup"


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

func CreateChannelGroupRepository(log logger.Logger) *repository.DBChannelGroupRepository {
	collectionName := getEnvOrDefault("MONGO_LIBRO_COLLECTION", DefaultChannelGroupCollectionName)
	repo, err := repository.CreateDBChannelGroupRepository(createMongoConfig(log), log, collectionName)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating bin repository")
		return nil
	}
	return repo
}



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

