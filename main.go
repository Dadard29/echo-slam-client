package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

func GetConnector() *EchoSlam {
  return Connector
}

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
  app.Bind(GetConnector())

  return app
}

func main() {
  if err := SetAccessor(); err != nil {
    Error(err)
  }

  SetConnector()

  app := getApp()

  if err := app.Run(); err != nil {
   Error(err)
  }
}
