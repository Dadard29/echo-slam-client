package main

import (
	"echo-slam-client/backend"
	"echo-slam-client/backend/log"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

var EchoSlam ConnectorInterface
var Accessor *Config
var ClientLogger *log.Logger


func setAccessor() error {
	h, err := homedir.Dir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(h, configFilePath)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// config file not exist
		Accessor, err = NewConfig(configFile)
		if err != nil {
			return err
		}
	} else {
		// config file exists
		Accessor, err = LoadConfig(configFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func setLogger() {
	ClientLogger = log.NewLogger()
}

func setConnector() {
	EchoSlam = &Connector{
		client: backend.NewApiClient(
			Accessor.Host, Accessor.JWT, Accessor.JWT != "", ClientLogger),
	}
	ClientLogger.Info(fmt.Sprintf("using connector with api %s", Accessor.Host))
}

func setConnectorMock() {
	EchoSlam = &ConnectorMock{}
	ClientLogger.Info(fmt.Sprintf("using MOCK connector"))
}

func setGlobals() {
	setLogger()
	if err := setAccessor(); err != nil {
		ClientLogger.Fatal(err)
	}

	setConnector()
}

func setGlobalsMock() {
	setLogger()
	if err := setAccessor(); err != nil {
		ClientLogger.Fatal(err)
	}

	setConnectorMock()
}

func unsetAccessorConnector() {
	Accessor = nil
	EchoSlam = nil
}


