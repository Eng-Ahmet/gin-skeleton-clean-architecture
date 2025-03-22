package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once          sync.Once
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

// initializeLoggers  initializes the loggers
// It is called only once when the first logger is requested
func initializeLoggers() {
	once.Do(func() {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Println("Failed to open log file, using default stderr:", err)
			file = os.Stderr // log to stderr if file is not available
		}

		infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// GetInfoLogger  returns the info logger
func GetInfoLogger() *log.Logger {
	initializeLoggers()
	return infoLogger
}

// GetWarningLogger  returns the warning logger
func GetWarningLogger() *log.Logger {
	initializeLoggers()
	return warningLogger
}

// GetErrorLogger  returns the error logger
func GetErrorLogger() *log.Logger {
	initializeLoggers()
	return errorLogger
}
