package helper

import (
	"os"
	"strconv"

	logger "gitlab.com/fravega-it/arquitectura/ec-golang-logger"
)

func GetEnvOrDefault(envName string, defaultValue string) string {
	if env := os.Getenv(envName); env == "" {
		return defaultValue
	} else {
		return env
	}
}

func GetEnvOrDefaultBool(envName string, defaultValue bool) bool {
	if env := os.Getenv(envName); env == "" {
		return defaultValue
	} else {
		return env == "true"
	}
}

func GetEnvOrDefaultInt(envName string, defaultValue int) int {
	if env := os.Getenv(envName); env == "" {
		return defaultValue
	} else {
		v, err := strconv.Atoi(env)
		if err != nil {
			logger.GetDefaultLogger().WithFields(logger.Fields{"envName": envName, "value": env}).Errorf("Invalid numeric value, using default (%v)", defaultValue)
			return defaultValue
		}
		return v
	}
}
