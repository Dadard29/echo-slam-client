package main

import (
	"echo-slam-client/backend/models"
	"fmt"
	"net/http"
)

const endpointProfile = "/profile"
const endpointProfilePublic = "/profile/public"
const endpointSignIn = "/signin"
const endpointAvatarList = "/avatar/list"
const endpointAvatar = "/avatar"
const endpointSignUp = "/signup"

func (e *Connector) GetProfile() (*models.Profile, error) {

	var profile = &models.ProfileResponse{}
	err := e.client.DoRequestProtectedJWT(endpointProfile, http.MethodGet, nil, profile)
	return &profile.Profile, err
}

func (e *Connector) GetProfilePublic(username string) (*models.ProfilePublic, error) {

	var profile = &models.ProfilePublicResponse{}

	params := map[string]string{
		"username": username,
	}
	err := e.client.DoRequestProtectedJWT(endpointProfilePublic, http.MethodGet, params, profile)
	return &profile.ProfilePublic, err
}

func (e *Connector) SignIn(username string, password string) error {

	var signIn = &models.SignInResponse{}

	err := e.client.DoRequestProtectedBasic(
		endpointSignIn, http.MethodGet, nil,
		username, password, signIn)

	if err == nil {
		Connect(username, signIn.JWT)
		return nil
	}

	return err
}

func (e *Connector) SignInCheck() error {

	var signInCheck = &models.SignInCheckResponse{}

	err := e.client.DoRequestProtectedJWT(
		endpointSignIn, http.MethodPost, nil, signInCheck)
	return err
}

func (e *Connector) LogOut() {
	Disconnect()
}

func (e *Connector) GetAvatarList() ([]string, error) {
	var avatarList = &models.AvatarListResponse{}
	err := e.client.DoRequest(
		endpointAvatarList, http.MethodGet, nil, avatarList)
	return avatarList.AvatarList, err
}

func (e *Connector) GetAvatarUrl(avatarName string) string {
	return fmt.Sprintf("%s%s?name=%s",
		e.client.Host(), endpointAvatar, avatarName)
}

func (e *Connector) IsUsernameUsed(username string) (bool, error) {
	var usernameCheckResponse = &models.UsernameCheckResponse{}
	params := map[string]string{
		"username": username,
	}
	err := e.client.DoRequest(
		endpointProfilePublic, http.MethodPost, params, usernameCheckResponse)
	return usernameCheckResponse.IsUsed, err
}

func (e *Connector) SignUp(
	recoverBy string, contact string, avatar string,
	username string, password string) error {

	var signUpResponse = &models.SignUpResponse{}
	params := map[string]string{
		"recover_by": recoverBy,
		"contact": contact,
		"avatar": avatar,
	}
	err := e.client.DoRequestProtectedBasic(
		endpointSignUp, http.MethodPost, params,
		username, password, signUpResponse)
	return err
}

func (e *Connector) SignUpConfirm(
	code string, username string, password string) error {

	var signUpConfirm = &models.ProfileResponse{}
	params := map[string]string{
		"code": code,
	}
	err := e.client.DoRequestProtectedBasic(
		endpointProfile, http.MethodPost, params,
		username, password, signUpConfirm)
	return err
}