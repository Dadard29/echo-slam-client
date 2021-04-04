package main

import (
	"echo-slam-client/backend/models"
	"net/http"
)

const endpointInfos = "/infos"
const endpointHealth = "/health/apis"

func (e *Connector) Infos() (*models.Infos, error) {
	var infos = &models.Infos{}
	err := e.doRequest(endpointInfos, http.MethodGet, nil, infos)
	return infos, err
}

func (e *Connector) HealthApis() (*models.HealthApis, error) {
	var health = &models.HealthApisResponse{}
	err := e.doRequest(endpointHealth, http.MethodGet, nil, health)
	return &health.Content, err
}


