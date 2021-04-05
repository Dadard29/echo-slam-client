package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var username = os.Getenv("USERNAME")
var password = os.Getenv("PASSWORD")

const avatarName = "avatar_aftershock"
const newUsernameTest = "test_3"
const newPasswordTest = "test_3test_3"

func TestConnectorProfilePublic(t *testing.T) {
	setGlobals()

	t.Run("avatar list", testConnectorGetAvatarList)
	t.Run("avatar URL", testConnectorGetAvatarUrl)
	t.Run("username check", testConnectorIsUsernameUsed)

	unsetAccessorConnector()
}

func TestConnectorProfileProtected(t *testing.T) {
	setGlobals()

	t.Run("signIn", testConnectorSignIn)
	t.Run("signIn check", testConnectorSignInCheck)
	t.Run("profile", testConnectorGetProfile)
	t.Run("profile public", testConnectorGetProfilePublic)
	t.Run("logOut", testConnectorLogOut)

	unsetAccessorConnector()
}

func testConnectorGetProfile(t *testing.T) {
	p, err := EchoSlam.GetProfile()
	if err != nil {
		t.Error(err)
	}

	t.Log(p.Username, p.Contact)
}

func testConnectorGetProfilePublic(t *testing.T) {
	username := "dadard"
	p, err := EchoSlam.GetProfilePublic(username)
	if err != nil {
		t.Error(err)
	}

	t.Log(p.Username)
}

func testConnectorSignIn(t *testing.T) {
	err := EchoSlam.SignIn(username, password)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, EchoSlam.Client().Connected())
}

func testConnectorSignInCheck(t *testing.T) {
	err := EchoSlam.SignInCheck()

	if EchoSlam.Client().Connected() {
		assert.True(t, err == nil, err)
	} else {
		assert.False(t, err == nil, err)
	}
}

func testConnectorLogOut(t *testing.T) {
	EchoSlam.LogOut()
	assert.False(t, EchoSlam.Client().Connected())
}

func testConnectorGetAvatarList(t *testing.T) {
	l, err := EchoSlam.GetAvatarList()
	if err != nil {
		t.Error(err)
	}

	for _, v := range l {
		t.Log(v)
	}
}

func testConnectorGetAvatarUrl(t *testing.T) {
	url := EchoSlam.GetAvatarUrl(avatarName)
	assert.NotEqual(t, "", url)
	t.Log(url)
}

func testConnectorIsUsernameUsed(t *testing.T) {
	check, err := EchoSlam.IsUsernameUsed(newUsernameTest)
	if err != nil {
		t.Error(err)
	}

	assert.False(t, check)

	check, err = EchoSlam.IsUsernameUsed(username)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, check)
}
