package logs

import (
	"errors"
	"log"
)

const (
	//Colors
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

type MyError struct {
	Error error
	Code  int
}

func LogFatal(msg string, code int) MyError {
	log.Fatal(string(red), msg, string(reset))
	return MyError{
		Error: errors.New(msg),
		Code:  code,
	}
}

func LogError(msg string, Code int) MyError {
	log.Println(string(red), msg, string(reset))
	return MyError{
		Error: errors.New(msg),
		Code:  Code,
	}
}

func LogWarning(msg string) {
	log.Println(string(yellow), "¡WARNING:", msg+"!", string(reset))
}

func Log(msg string) {
	log.Println(msg)
}

func LogSuccess(msg string) {
	log.Println(string(green), msg, string(reset))
}
