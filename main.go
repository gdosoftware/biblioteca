package main

import (
    
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	"github.com/gdosoftware/biblioteca/infraestructura/api"
    "github.com/gdosoftware/biblioteca/infraestructura/module"
    "github.com/gdosoftware/biblioteca/infraestructura/health"
	"github.com/gdosoftware/biblioteca/infraestructura"
)

const (
	// Application
	AppName    = "Medios de pago"
	AppVersion = "0.1.0"
)

func main() {

	log := logger.GetDefaultLogger()
	log.Infof("############ Starting Medio de pago Service %v ############", AppVersion)

	info := api.NewBuildInfo(AppName, AppVersion)

	err := module.NewSourceFactory()
	if err != nil {
		log.Fatal(err)
	}

    // Repositorios
    libroRepository := module.CreateLibroRepository(log)
  //  socioRepository := module.CreateSocioRepository(log)
   // prestamoRepository := module.CreatePrestamoRepository(log)
	
	/*	jwtDecoder := module.CreateJwtDecoder()
		jwtTokenTask := module.GetTokenTask()*/


	
	// Controller & Server
	log.Debug("Creating Server")
	controllers := module.MakeControllers(libroRepository)
	sensors := []health.Sensor{libroRepository}
	appServer := server.NewRestServer(info, controllers, sensors)
	defer appServer.Done()
	appServer.Run()
}
