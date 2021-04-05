package main

import (
	"echo-slam-client/backend/models"
)

type ConnectorMock struct {

}

func (c *ConnectorMock) Connected() bool {
	return false
}

func (c *ConnectorMock) Connect(JWT string) {
	return
}

func (c *ConnectorMock) Disconnect() {
	return
}


func (c *ConnectorMock) Infos() (*models.Infos, error) {
	return &models.Infos{
		ContactEmail: "email",
		Description:  "desc",
		License:      "license",
		LicenseUrl:   "wtfpl.net",
		Title:        "title",
		Version:      "0.0.1",
	}, nil
}

func (c *ConnectorMock) HealthApis() (*models.HealthApis, error) {
	return &models.HealthApis{
		List: []models.ApiCheck{
			{APIName: "api", Check:   true},
			{APIName: "api", Check:   true},
			{APIName: "api", Check:   true},
		},
	}, nil
}

func (c *ConnectorMock) GetProfile() (*models.Profile, error) {
	panic("implement me")
}

func (c *ConnectorMock) GetProfilePublic(username string) (*models.ProfilePublic, error) {
	panic("implement me")
}

func (c *ConnectorMock) SignIn(username string, password string) error {
	panic("implement me")
}

func (c *ConnectorMock) SignInCheck() error {
	panic("implement me")
}

func (c *ConnectorMock) LogOut() {
	panic("implement me")
}

func (c *ConnectorMock) GetAvatarList() ([]string, error) {
	panic("implement me")
}

func (c *ConnectorMock) GetAvatarUrl(avatarName string) string {
	panic("implement me")
}

func (c *ConnectorMock) IsUsernameUsed(username string) (bool, error) {
	panic("implement me")
}

func (c *ConnectorMock) SignUp(recoverBy string, contact string, avatar string, username string, password string) error {
	panic("implement me")
}

func (c *ConnectorMock) SignUpConfirm(code string, username string, password string) error {
	panic("implement me")
}


