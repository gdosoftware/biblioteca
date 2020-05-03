package main

import (
    
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/infrastructure/http"
    "github.com/gdosoftware/biblioteca/infrastructure/module"
    "github.com/gdosoftware/biblioteca/infrastructure/health"
	"github.com/gdosoftware/biblioteca/infrastructure"
)

const (
	// Application
	AppName    = "Agrupaciones"
	AppVersion = "0.1.0"
)

func main() {

	log := logger.GetDefaultLogger()
	log.Infof("############ Starting Agrupaciones Service %v ############", AppVersion)

	info := http.NewBuildInfo(AppName, AppVersion)

	err := module.NewSourceFactory()
	if err != nil {
		log.Fatal(err)
	}

    // Repositorios
    channelGroupRepository := module.CreateChannelGroupRepository(log)
	
	/*	jwtDecoder := module.CreateJwtDecoder()
		jwtTokenTask := module.GetTokenTask()*/


	
	// Controller & Server
	log.Debug("Creating Server")
	controllers := module.MakeControllers(channelGroupRepository)
	sensors := []health.Sensor{channelGroupRepository}
	appServer := server.NewRestServer(info, controllers, sensors)
	defer appServer.Done()
	appServer.Run()
}
