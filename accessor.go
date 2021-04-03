package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
)

const defaultHost = "https://api.echo-slam.com"
const configFilePath = ".config/echo-slam/echo-slam.json"

// access the user config
type Accessor struct {
	path string

	Host     string
	Username string
	JWT      string
}

func (a *Accessor) Save() error {
	data, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(a.path, data, 0644)
}

func LoadConfig(configFile string) (*Accessor, error) {
	Info(fmt.Sprintf("loading existing config at %s", configFile))

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var accessor = &Accessor{}
	err = json.Unmarshal(data, accessor)
	if err != nil {
		return nil, err
	}

	accessor.path = configFile

	return accessor, nil
}

func NewConfig(configFile string) (*Accessor, error) {
	Info(fmt.Sprintf("setting up new config at %s", configFile))

	newAccessor := &Accessor{
		path:     configFile,
		Host:     defaultHost,
		Username: "",
		JWT:      "",
	}

	fileDir := filepath.Dir(configFile)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		if err := os.Mkdir(fileDir, 0755); err != nil {
			return nil, err
		}
	}

	if err := newAccessor.Save(); err != nil {
		return nil, err
	}
	return newAccessor, nil
}

func SetAccessor() error {
	h, err := homedir.Dir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(h, configFilePath)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// config file not exist
		Config, err = NewConfig(configFile)
		if err != nil {
			return err
		}
	} else {
		// config file exists
		Config, err = LoadConfig(configFile)
		if err != nil {
			return err
		}
	}

	return nil
}
