package logging

import (
	"log"
	"os"
)

var Logger = log.New(os.Stderr, "", log.LstdFlags)
var debugEnabled = false

func EnableDebug() {
	debugEnabled = true
}

func Info(msg string) {
	Logger.Println("[INFO] " + msg)
}

func Debug(msg string) {
	if debugEnabled {
		Logger.Println("[DEBUG] " + msg)
	}
}

func Error(msg string) {
	Logger.Println("[ERROR] " + msg)
}

