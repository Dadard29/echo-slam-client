package log

import (
	"github.com/fatih/color"
	"github.com/wailsapp/wails"
	"log"
)

const levelInfo = "INFO"
const levelError = "ERROR"
const levelLoading = "LOADING"
const levelDebug = "DEBUG"

type LoggerMessage struct {
	Msg string `json:"msg"`
	Level string `json:"level"`
}

type Logger struct {
	msgList []LoggerMessage
	runtime *wails.Runtime
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) WailsInit(runtime *wails.Runtime) error {
	l.runtime = runtime
	return nil
}


func (l *Logger) GetLastMsg() LoggerMessage {
	if len(l.msgList) > 0 {
		return l.msgList[len(l.msgList) - 1]
	} else {
		return LoggerMessage{"no message", levelDebug}
	}
}

func (l *Logger) log(msg string, fgColor color.Attribute, level string) {
	color.Set(fgColor)
	log.Println(msg)
	color.Unset()

	l.msgList = append(l.msgList, LoggerMessage{
		msg, level,
	})

	if l.runtime != nil {
		l.runtime.Events.Emit("log")
	}
}

func (l *Logger) Info(msg string) {
	l.log(msg, color.FgGreen, levelInfo)
}

func (l *Logger) Error(msg string) {
	l.log(msg, color.FgRed, levelError)
}

func (l *Logger) Debug(msg string) {
	l.log(msg, color.FgWhite, levelDebug)
}

func (l *Logger) Loading(msg string) {
	l.log(msg, color.FgWhite, levelLoading)
}

func (l *Logger) Fatal(err error) {
	color.Set(color.FgHiRed)
	log.Fatalln(err.Error())
}
