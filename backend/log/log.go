package log

import (
	"github.com/fatih/color"
	"log"
)

func Info(msg string) {
	color.Set(color.FgGreen)
	log.Println(msg)
	color.Unset()
}

func Debug(msg string) {
	log.Println(msg)
}

func Error(err error) {
	color.Set(color.FgRed)
	log.Println(err.Error())
	color.Unset()
}

func Fatal(err error) {
	color.Set(color.FgHiRed)
	log.Fatalln(err.Error())
}
