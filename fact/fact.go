package fact

import "fmt"

// we have a very commen logger
type Logger interface {
	Log(message string)
}

// we need to implement certain logging for example
// 1. console.logging , 2. file logging

type ConsoleLogger struct{}

func (cl *ConsoleLogger) Log(message string) {
	fmt.Println("logging on console")
}

type FileLogger struct{}

func (fl *FileLogger) Log(message string) {
	fmt.Println("logging in file")
}

type LoggerFactory interface {
	CreateLogger() Logger
}

type FileLoggerFactory struct{}

func (flf *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

type ConsoleLoggerFactory struct{}

func (clf *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
