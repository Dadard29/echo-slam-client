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

func (e *ConnectorImplement) GetProfile() (*models.Profile, error) {

	var profile = &models.ProfileResponse{}
	err := e.Client().DoRequestProtectedJWT(endpointProfile, http.MethodGet, nil, profile)
	return &profile.Profile, err
}

func (e *ConnectorImplement) GetProfilePublic(username string) (*models.ProfilePublic, error) {

	var profile = &models.ProfilePublicResponse{}

	params := map[string]string{
		"username": username,
	}
	err := e.Client().DoRequestProtectedJWT(endpointProfilePublic, http.MethodGet, params, profile)
	return &profile.ProfilePublic, err
}

func (e *ConnectorImplement) SignIn(username string, password string) error {

	var signIn = &models.SignInResponse{}

	err := e.Client().DoRequestProtectedBasic(
		endpointSignIn, http.MethodGet, nil,
		username, password, signIn)

	if err == nil {
		Connect(username, signIn.JWT)
		return nil
	}

	return err
}

func (e *ConnectorImplement) SignInCheck() error {

	var signInCheck = &models.SignInCheckResponse{}

	err := e.Client().DoRequestProtectedJWT(
		endpointSignIn, http.MethodPost, nil, signInCheck)
	return err
}

func (e *ConnectorImplement) LogOut() {
	Disconnect()
}

func (e *ConnectorImplement) GetAvatarList() ([]string, error) {
	var avatarList = &models.AvatarListResponse{}
	err := e.Client().DoRequest(
		endpointAvatarList, http.MethodGet, nil, avatarList)
	return avatarList.AvatarList, err
}

func (e *ConnectorImplement) GetAvatarUrl(avatarName string) string {
	return fmt.Sprintf("%s%s?name=%s",
		e.Client().Host(), endpointAvatar, avatarName)
}

func (e *ConnectorImplement) IsUsernameUsed(username string) (bool, error) {
	var usernameCheckResponse = &models.UsernameCheckResponse{}
	params := map[string]string{
		"username": username,
	}
	err := e.Client().DoRequest(
		endpointProfilePublic, http.MethodPost, params, usernameCheckResponse)
	return usernameCheckResponse.IsUsed, err
}

func (e *ConnectorImplement) SignUp(
	recoverBy string, contact string, avatar string,
	username string, password string) error {

	var signUpResponse = &models.SignUpResponse{}
	params := map[string]string{
		"recover_by": recoverBy,
		"contact": contact,
		"avatar": avatar,
	}
	err := e.Client().DoRequestProtectedBasic(
		endpointSignUp, http.MethodPost, params,
		username, password, signUpResponse)
	return err
}

func (e *ConnectorImplement) SignUpConfirm(
	code string, username string, password string) error {

	var signUpConfirm = &models.ProfileResponse{}
	params := map[string]string{
		"code": code,
	}
	err := e.Client().DoRequestProtectedBasic(
		endpointProfile, http.MethodPost, params,
		username, password, signUpConfirm)
	return err
}