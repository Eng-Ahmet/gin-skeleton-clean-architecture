package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func InitLogger(infoLogPath, errorLogPath string) {
	var infoLogOutput, errorLogOutput *os.File
	var err error

	if infoLogPath != "" {
		infoLogOutput, err = os.OpenFile(infoLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open info log file: %v", err)
		}
	} else {
		infoLogOutput = os.Stdout
	}

	if errorLogPath != "" {
		errorLogOutput, err = os.OpenFile(errorLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open error log file: %v", err)
		}
	} else {
		errorLogOutput = os.Stderr
	}

	infoLogger = log.New(infoLogOutput, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorLogOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	infoLogger.Println(message)
}

func Error(message string) {
	errorLogger.Println(message)
}
