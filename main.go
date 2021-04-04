package main

import (
  "echo-slam-client/backend/log"
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
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

  app.Bind(Accessor)
  app.Bind(EchoSlam)

  return app
}

func main() {
  setAccessorConnector()

  app := getApp()

  if err := app.Run(); err != nil {
   log.Error(err)
  }
}
