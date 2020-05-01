package api

import (
	"net/http"
	"runtime"

	"github.com/ant0ine/go-json-rest/rest"
	"gitlab.com/fravega-it/adn/ipos/medios-de-pago/backend-medios-de-pago/infrastructure/health"
)

type BuildInfo struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
}

func NewBuildInfo(appName string, appVersion string) *BuildInfo {
	return &BuildInfo{appName, appVersion, runtime.Version()}
}

type SysStatusAPI struct {
	buildInfo *BuildInfo
	sensors   []health.Sensor
	statusMw  *rest.StatusMiddleware
}

func NewSystemAPI(buildInfo *BuildInfo, statusMw *rest.StatusMiddleware, sensors ...health.Sensor) *SysStatusAPI {
	return &SysStatusAPI{buildInfo: buildInfo, sensors: sensors, statusMw: statusMw}
}

func (c *SysStatusAPI) Info(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(c.buildInfo)
}

// Report the status of all the sub-systems sensors
func (c *SysStatusAPI) Status(w rest.ResponseWriter, r *rest.Request) {
	if len(c.sensors) > 0 {
		var stats = make(map[string]interface{})

		for _, sensor := range c.sensors {
			stats[sensor.Name()] = sensor.Stats()
		}
		w.WriteJson(stats)
	}
}

// Reports Server statistics values
func (c *SysStatusAPI) Stats(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(c.statusMw.GetStatus())
}

// Report the status of all sensors. If any of then on error, report as service unavailable
func (c *SysStatusAPI) Health(w rest.ResponseWriter, r *rest.Request) {
	if len(c.sensors) > 0 {
		var errors = make(map[string]string)

		for _, sensor := range c.sensors {
			if err := sensor.Health(); err != nil {
				errors[sensor.Name()] = err.Error()
			}
		}

		if len(errors) > 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.WriteJson(errors)
		} else {
			w.WriteJson("true")
		}
	}
}
