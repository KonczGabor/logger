package logger

import "fmt"

func Info(message string, data interface{}) {
	fmt.Println("Log info: ", message, data)
}

func Fatal(message, data interface{}) {
	fmt.Println("Log Fatal: ", message, data)
}
