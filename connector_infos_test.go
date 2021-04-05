package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectorInfos(t *testing.T) {
	setGlobals()

	t.Run("infos", testConnectorInfos)
	t.Run("health", testConnectorHealthApis)
}

func testConnectorInfos(t *testing.T) {
	infos, err := EchoSlam.Infos()
	if err != nil {
		t.Error(err)
	}
	t.Log(infos.Description)
}

func testConnectorHealthApis(t *testing.T) {
	health, err := EchoSlam.HealthApis()
	if err != nil {
		t.Error(err)
	}
	for _, v := range health.List {
		t.Logf("checking %s...", v.APIName)
		assert.True(t,  v.Check)
	}
}
