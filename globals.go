package main

import (
	"echo-slam-client/backend/log"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"net/http"
	"os"
	"path/filepath"
)

var EchoSlam *Connector
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
	EchoSlam = &Connector{
		client:    &http.Client{},
		Host:      Accessor.Host,
		JWT:       Accessor.JWT,
		Connected: Accessor.JWT != "",
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


func Connect(JWT string) {
	Accessor.JWT = JWT
	if err := Accessor.Save(); err != nil {
		log.Error(err)
	}

	EchoSlam.JWT = JWT
	EchoSlam.Connected = true
}

func Disconnect() {
	Accessor.JWT = ""
	if err := Accessor.Save(); err != nil {
		log.Error(err)
	}

	EchoSlam.JWT = ""
	EchoSlam.Connected = false
}
