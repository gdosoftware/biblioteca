package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
	
	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
	api2 "github.com/gdosoftware/biblioteca/infraestructura/http"
	"github.com/gdosoftware/biblioteca/infraestructura/health"
	"github.com/gdosoftware/biblioteca/infraestructura/helper"
	"github.com/gdosoftware/biblioteca/infraestructura/controllers"
)

const (
	DefaultPortValue    = 9000
	PortEnvironmentName = "SERVER_PORT"
)

type Controller interface {
	Routes() []*rest.Route
}

// A Rest Server that listen for the HTTP requests
type RestServer struct {
	root string
	port int
	api  *rest.Api
}

func NewRestServer(info *api2.BuildInfo, controllers []Controller, sensors []health.Sensor) *RestServer {

	apiName := "/biblioteca/api/v1"
	port, err := strconv.Atoi(helper.GetEnvOrDefault(PortEnvironmentName, strconv.Itoa(DefaultPortValue)))
	if err != nil {
		port = DefaultPortValue
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": err, "usingPort": DefaultPortValue}).Warn("Couldn't not use specified port, going to use default one: 9000")
	}

	logger.GetDefaultLogger().Debug("####### Configuring Medio de Pago Service Module #######")

	api := rest.NewApi()

	stack := []rest.Middleware{
		&helper.AccessLogMiddleware{Logger: logger.GetDefaultLogger()},
		&rest.TimerMiddleware{},
		&rest.RecorderMiddleware{},
		&rest.RecoverMiddleware{
			EnableResponseStackTrace: true,
		},
		&rest.JsonIndentMiddleware{},
		&rest.ContentTypeCheckerMiddleware{},
		&rest.GzipMiddleware{},
	}

	statusMw := &rest.StatusMiddleware{}

	api.Use(statusMw)

	api.Use(stack...)

	router, err := makeRouterForControllers(controllers)

	if err != nil {
		logger.GetDefaultLogger().Fatal(err)
		panic(err)
	}

	api.SetApp(router)

	http.Handle(apiName+"/", http.StripPrefix(apiName, api.MakeHandler()))

	err = systemApi(info, statusMw, sensors)

	if err != nil {
		logger.GetDefaultLogger().Fatal(err)
		panic(err)
	}

	return &RestServer{port: port, root: apiName, api: api}
}

func makeRouterForControllers(controllers []Controller) (rest.App, error) {
	var routes = make([]*rest.Route, 0)
	// Add routes for all registered controllers
	for _, controller := range controllers {
		routes = append(routes, controller.Routes()...)
	}
	// Add the routes of the system controller
	return rest.MakeRouter(routes...)
}

// Adds the endpoints for the System controller
func systemApi(info *api2.BuildInfo, statusMw *rest.StatusMiddleware, sensors []health.Sensor) error {
	systemApi := api2.NewSystemAPI(info, statusMw, sensors...)
	systemStat := controllers.NewSystemController(systemApi)

	api := rest.NewApi()

	router, err := rest.MakeRouter(systemStat.Routes()...)
	if err != nil {
		return err
	}

	api.SetApp(router)

	http.Handle("/biblioteca/system/", http.StripPrefix("/biblioteca/system", api.MakeHandler()))

	return nil
}

// Start the server and keep it running until it is stopped
func (s *RestServer) Run() {
	logger.GetDefaultLogger().Info("Listening on :", s.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil); err != nil {
		panic(err)
	}
}

func (s *RestServer) Done() {
	logger.GetDefaultLogger().Info("Done everything")
	if x := recover(); x != nil {
		logger.GetDefaultLogger().WithFields(logger.Fields{"error": x}).Error("Run time panic")
	}
}
