package singleton

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	file *os.File
	*log.Logger
}

var (
	instance *Logger
	once     sync.Once
)

func GetInstance() *Logger {
	once.Do(func() {
		// os.O_APPEND and os.O_CREATE and os.O_WRONLY are integer constant and here we by doing OR between we are setting
		// the flag in such a way that the resultant flag contains all these options set to the file
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("failed to open/create the log file: %v", err.Error())
		}
		instance = &Logger{
			file:   file,
			Logger: log.New(file, "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})
	return instance
}

func (l *Logger) Close() {
	l.file.Close()
}
