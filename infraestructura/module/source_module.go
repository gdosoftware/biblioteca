package module

import (
	"fmt"
	"os"
	"strings"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/infraestructura/repository/mongo"
	"github.com/gdosoftware/biblioteca/infraestructura/api"

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
}


func CreateJwtDecoder() *api.JwtDecoder {
	return action.CreateJwtDecoder(os.Getenv(JWTPublicKeyEnv))
}

func GetTokenTask() string {
	return os.Getenv(JWTTokenTask)
}

/*
func (s *SourceFactory) CreateComponedItemRepository() *source.ComponedRepository {
	katalogRepo := item.CreateKatalogRepository(os.Getenv(KatalogServiceEndpointEnv))
	priceRepo := item.CreatePriceRepository(os.Getenv(PriceServiceEndpointEnv), os.Getenv(PriceServiceNotificationEndpointEnv), os.Getenv(PriceServiceBotNotificationEndpointEnv))
	stockRepo := item.CreateStockRepository(os.Getenv(StockServiceEndpointEnv),
		os.Getenv(StockServiceNotificationEndpointEnv))
	mediaRepo := item.CreateMediaRepository(os.Getenv(MediaServiceEndpointEnv))
	creator := item.CreateItemCreator(mediaRepo, katalogRepo)
	logisticsRepository := s.CreateLogisticsRepository()
	repo, err := source.CreateItemRepository(katalogRepo, priceRepo, stockRepo, creator, logisticsRepository.GetDefaultWarehouse)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating item repository")
		return nil
	}
	return repo
}
*/
/*
func (s *SourceFactory) CreatePaymentMethodBinRepository() *source.configBin {
	collectionName := helper.GetEnvOrDefault(MongoDBSKUCollectionName, DefaultSkuCollectionName)
	repo, err := source.CreateDBSkuRepository(s.createMongoConfig(), collectionName)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating sku repository")
		return nil
	}
	return repo
}
*/

/*
func (s *SourceFactory) CreateMongoBasedSkuRepository() *source.DBSkuRepository {
	collectionName := helper.GetEnvOrDefault(MongoDBSKUCollectionName, DefaultSkuCollectionName)
	repo, err := source.CreateDBSkuRepository(s.createMongoConfig(), collectionName)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating sku repository")
		return nil
	}
	return repo
}

func (s *SourceFactory) CreateJobRepository() *source.MongoJobRepository {
	repo, err := source.CreateMongoJobRepository(s.createMongoConfig())
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating task repository")
		return nil
	}
	return repo
}

func (s *SourceFactory) CreateMongoTermsRepository() *source.DBTermsRepository {
	repo, err := source.CreateDBDBTermsRepository(s.createMongoConfig())
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating terms repository")
		return nil
	}
	return repo
}

func (s *SourceFactory) createMongoConfig() *mgo.DialInfo {
	databaseName, dataBaseAddr := os.Getenv(MongoDBName), os.Getenv(MongoDBAddress)
	if databaseName == "" || dataBaseAddr == "" {
		logger.GetDefaultLogger().Fatal("Database name/address is mandatory")
		return nil
	}
	return &mgo.DialInfo{
		Addrs:    strings.Split(dataBaseAddr, ","),
		Database: databaseName,
	}
}

func (s *SourceFactory) CreateComponedItemRepository() *source.ComponedItemRepository {
	katalogRepo := item.CreateKatalogRepository(os.Getenv(KatalogServiceEndpointEnv))
	priceRepo := item.CreatePriceRepository(os.Getenv(PriceServiceEndpointEnv), os.Getenv(PriceServiceNotificationEndpointEnv), os.Getenv(PriceServiceBotNotificationEndpointEnv))
	stockRepo := item.CreateStockRepository(os.Getenv(StockServiceEndpointEnv),
		os.Getenv(StockServiceNotificationEndpointEnv))
	mediaRepo := item.CreateMediaRepository(os.Getenv(MediaServiceEndpointEnv))
	creator := item.CreateItemCreator(mediaRepo, katalogRepo)
	logisticsRepository := s.CreateLogisticsRepository()
	repo, err := source.CreateItemRepository(katalogRepo, priceRepo, stockRepo, creator, logisticsRepository.GetDefaultWarehouse)
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating item repository")
		return nil
	}
	return repo
}

func (s *SourceFactory) CreateSellerRepository() *source.SellerRepository {
	return source.CreateSellerRepository(os.Getenv(SellerServiceEndpointEnv), helper.GetEnvOrDefaultInt(SellerCacheRefreshInterval, 30))
}

func (s *SourceFactory) CreateJwtDecoder() *action.JwtDecoder {
	return action.CreateJwtDecoder(os.Getenv(JWTPublicKeyEnv))
}

func (s *SourceFactory) CreateCategoryRepository() *source.CategoryRepository {
	return source.CreateCategoryRepository(os.Getenv(KatalogServiceEndpointEnv), os.Getenv(AttributeServiceEndpointEnv))
}

func (s *SourceFactory) CreateLogisticsRepository() *source.LogisticsRepository {
	return source.CreateLogisticsRepository(os.Getenv(LogisticsServiceEndpointEnv))
}

func (s *SourceFactory) GetTokenTask() string {
	return os.Getenv(JWTTokenTask)
}

*/

func validateEnvironmentVariables() []string {
	var missingEnvVars []string
	for _, v := range mandatoryEnvironmentVariables {
		if strings.TrimSpace(os.Getenv(v)) == "" {
			missingEnvVars = append(missingEnvVars, v)
		}
	}
	return missingEnvVars
}

/*
func (s *SourceFactory) CreateSmtpMailer() *source.SmtpMailer {
	repo, err := source.NewSmtpMailer(s.createSmtpMailerConfig())
	if err != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err}).Fatal("Error while creating smtp mailer")
		return nil
	}
	return repo
}

func (s *SourceFactory) createSmtpMailerConfig() *source.SmtpServerConfig {
	smtpHostEnv, smtpPortEnv := os.Getenv(smtpHost), os.Getenv(smtpPort)
	if smtpHostEnv == "" || smtpPortEnv == "" {
		logger.GetDefaultLogger().Fatal("Smtp address/port is mandatory")
		return nil
	}
	return &source.SmtpServerConfig{
		Host:     smtpHostEnv,
		Port:     smtpPortEnv,
		Username: helper.GetEnvOrDefault(smtpUsername, ""),
		Password: helper.GetEnvOrDefault(smtpPassword, ""),
		Insecure: helper.GetEnvOrDefault(smtpAcceptInsecure, "false") == "true",
		From:     helper.GetEnvOrDefault(smtpFrom, "no-responder@fravega.com"),
	}
}
*/
