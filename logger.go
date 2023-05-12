package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Log info: ", data)
}
