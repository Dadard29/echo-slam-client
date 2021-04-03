package main

import (
	"fmt"
)

type EchoSlam struct {
	Host string
}

func SetConnector() {
	Connector = &EchoSlam{
		Host: Config.Host,
	}
	Info(fmt.Sprintf("setting up connector with api %s", Connector.Host))
}
