package logger

import (
	"fmt"
	"github.com/marcoshuck/todogo/errors"
	"log"
	"os"
	"sync"
)

// Logger defines an application logger
type Logger struct {
	*log.Logger
}


// instance is the logger instance
var instance *Logger

// once is being used to set up the singleton.
var once sync.Once

// New returns a new Logger instance
func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[LOG]", log.LstdFlags),
	}
}

// instantiate initializes the instance package variable
func instantiate() {
	instance = New()
}

// GetInstance returns the current Logger's instance.
// It uses the singleton pattern to always return the same instance.
func GetInstance() *Logger {
	once.Do(instantiate)
	return instance
}

// Log prints a log message.
func (l *Logger) Log(message string) {
	instance.SetPrefix("[LOG] ")
	instance.Println(message)
}

// Assert prints the result of a boolean condition.
func (l *Logger) Assert(condition bool) {
	instance.SetPrefix("[ASSERT] ")
	instance.Printf("%v", condition)
}

// Debug prints a debug message.
func (l *Logger) Debug(message string) {
	instance.SetPrefix("[DEBUG] ")
	instance.Println(message)
}

// Error prints an error message.
func (l *Logger) Error(err errors.Error) {
	instance.SetPrefix("[ERROR] ")
	instance.Println(fmt.Sprintf("[Error] Status: %d | Code: %d | Message: %s", err.Status, err.Code, err.Message))
	instance.Println(fmt.Sprintf("Stack trace: %v", err.Base))
}

// Info prints an information message.
func (l *Logger) Info(message string) {
	instance.SetPrefix("[INFO] ")
	instance.Println(message)
}

// Warn prints a warning message.
func (l *Logger) Warn(message string) {
	instance.SetPrefix("[WARNING] ")
	instance.Println(message)
}