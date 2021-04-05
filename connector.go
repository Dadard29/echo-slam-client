package main

import (
	"echo-slam-client/backend"
	"echo-slam-client/backend/models"
)

type ConnectorInterface interface {
	Connected() bool
	Connect(JWT string)
	Disconnect()

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

func (e *Connector) Connect(JWT string) {
	e.client.Connect(JWT)
}

func (e *Connector) Disconnect() {
	e.client.Disconnect()
}

func (e *Connector) Connected() bool {
	return e.client.Connected()
}


