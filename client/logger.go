package client

import "fmt"

func Error(message string, v interface{}) {
	fmt.Printf("[Error]\t%s => '%v' incorrect\n", message, v)
}

func Info(message string, v interface{}) {
	fmt.Printf("[Info]\t%s\t%v\n", message, v)
}