package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const defaultHost = "https://api.echo-slam.com"
const configFilePath = ".config/echo-slam/echo-slam.json"

// access the user config
type Config struct {
	path string

	Host     string `json:"host"`
	Username string `json:"username"`
	JWT      string `json:"JWT"`
}

func (a *Config) Self() *Config {
	return a
}

func (a *Config) Save() error {
	data, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(a.path, data, 0644)
}

func LoadConfig(configFile string) (*Config, error) {
	ClientLogger.Info(fmt.Sprintf("loading existing config at %s", configFile))

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var accessor = &Config{}
	err = json.Unmarshal(data, accessor)
	if err != nil {
		return nil, err
	}

	accessor.path = configFile

	return accessor, nil
}

func NewConfig(configFile string) (*Config, error) {
	ClientLogger.Info(fmt.Sprintf("using new config at %s", configFile))

	newAccessor := &Config{
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
