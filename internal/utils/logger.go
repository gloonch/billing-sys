package utils

import (
	"log"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func LogInfo(category, message, methodName string) {
	log.Printf("%s[INFO] [%s] (%s) %s %s", Cyan, category, methodName, Reset, message)
}

func LogError(category, message, methodName string) {
	log.Printf("%s[ERROR] [%s] (%s) %s %s", Red, category, methodName, Reset, message)
}

func LogSuccess(category, message, methodName string) {
	log.Printf("%s[SUCCESS] [%s] (%s) %s %s", Green, category, methodName, Reset, message)
}
