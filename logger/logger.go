package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
)

// Init initializes the logger
func Init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an info message
func Info(format string, args ...interface{}) {
	if infoLogger != nil {
		infoLogger.Printf(format, args...)
	} else {
		fmt.Printf("INFO: "+format+"\n", args...)
	}
}

// Error logs an error message
func Error(format string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Printf(format, args...)
	} else {
		fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	}
}

// Debug logs a debug message
func Debug(format string, args ...interface{}) {
	// Only log debug messages if DEBUG environment variable is set
	if os.Getenv("DEBUG") == "true" {
		if debugLogger != nil {
			debugLogger.Printf(format, args...)
		} else {
			fmt.Printf("DEBUG: "+format+"\n", args...)
		}
	}
}

// Fatal logs a fatal error and exits
func Fatal(format string, args ...interface{}) {
	if errorLogger != nil {
		errorLogger.Fatalf(format, args...)
	} else {
		fmt.Fprintf(os.Stderr, "FATAL: "+format+"\n", args...)
		os.Exit(1)
	}
}
