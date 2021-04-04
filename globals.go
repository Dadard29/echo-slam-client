package main

import (
	"echo-slam-client/backend"
	"echo-slam-client/backend/log"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

var EchoSlam Connector
var Accessor *Config


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

func setConnector() {
	EchoSlam = &ConnectorImplement{
		client: backend.NewApiClient(Accessor.Host, Accessor.JWT, Accessor.JWT != ""),
	}
	log.Info(fmt.Sprintf("setting up connector with api %s", Accessor.Host))
}

func setAccessorConnector() {
	if err := setAccessor(); err != nil {
		log.Fatal(err)
	}

	setConnector()
}

func unsetAccessorConnector() {
	Accessor = nil
	EchoSlam = nil
}


func Connect(username string, JWT string) {
	Accessor.JWT = JWT
	Accessor.Username = username
	if err := Accessor.Save(); err != nil {
		log.Error(err)
	}

	EchoSlam.Client().Connect(JWT)
}

func Disconnect() {
	Accessor.JWT = ""
	Accessor.Username = ""
	if err := Accessor.Save(); err != nil {
		log.Error(err)
	}

	EchoSlam.Client().Disconnect()
}
