package main

import (
	"echo-slam-client/backend/models"
	"net/http"
)

const endpointInfos = "/infos"
const endpointHealth = "/health/apis"

func (e *ConnectorImplement) Infos() (*models.Infos, error) {
	var infos = &models.Infos{}
	err := e.Client().DoRequest(endpointInfos, http.MethodGet, nil, infos)
	return infos, err
}

func (e *ConnectorImplement) HealthApis() (*models.HealthApis, error) {
	var health = &models.HealthApisResponse{}
	err := e.Client().DoRequest(endpointHealth, http.MethodGet, nil, health)
	return &health.Content, err
}


