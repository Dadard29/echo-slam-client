package main

import (
	"echo-slam-client/backend"
	"echo-slam-client/backend/models"
)

type ConnectorInterface interface {
	Client() *backend.ApiClient

	Infos() (*models.Infos, error)
	HealthApis() (*models.HealthApis, error)

	GetProfile() (*models.Profile, error)
	GetProfilePublic(username string) (*models.ProfilePublic, error)
	SignIn(username string, password string) error
	SignInCheck() error
	LogOut()
	GetAvatarList() ([]string, error)
	GetAvatarUrl(avatarName string) string
	IsUsernameUsed(username string) (bool, error)
	SignUp(recoverBy string, contact string, avatar string,
		username string, password string) error
	SignUpConfirm(code string, username string, password string) error
}

type Connector struct {
	client *backend.ApiClient
}

func (e *Connector) Client() *backend.ApiClient {
	return e.client
}


