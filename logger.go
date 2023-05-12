package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Log info: ", data)
}

func Fatal(data interface{}) {
	fmt.Println("Log Fatal: ", data)
}
