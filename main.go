package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
  "os"
)

func getApp() *wails.App {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1920,
    Height: 1080,
    Title:  "echo-slam-client",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })

  app.Bind(ClientLogger)
  app.Bind(Accessor)
  app.Bind(EchoSlam)

  return app
}

func main() {
  if os.Getenv("DEBUG") == "1" {
    setGlobalsMock()
  } else {
    setGlobals()
  }

  app := getApp()

  if err := app.Run(); err != nil {
    ClientLogger.Error(err.Error())
  }
}
