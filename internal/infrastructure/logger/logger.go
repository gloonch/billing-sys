package logger

import "log"

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l Logger) Info(msg string) {
	log.Println("[INFO] " + msg)
}
