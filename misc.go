package main

func Connect(username string, JWT string) {
	Accessor.JWT = JWT
	Accessor.Username = username
	if err := Accessor.Save(); err != nil {
		ClientLogger.Error(err.Error())
	}

	EchoSlam.Connect(JWT)
}

func Disconnect() {
	Accessor.JWT = ""
	Accessor.Username = ""
	if err := Accessor.Save(); err != nil {
		ClientLogger.Error(err.Error())
	}

	EchoSlam.Disconnect()
}
