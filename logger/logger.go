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

// FormatErrorMessage converts an err instance to an error message
func (l *Logger) FormatErrorMessage(err errors.Error) string {
	return fmt.Sprintf("Code: %d | Status: %d | Message: %s", err.Code, err.Status, err.Message)
}

// Log prints a log message.
func (l *Logger) Log(a ...interface{}) {
	instance.SetPrefix("[LOG] ")
	instance.Println(a...)
}

// Assert prints the result of a boolean condition.
func (l *Logger) Assert(condition bool) {
	instance.SetPrefix("[ASSERT] ")
	instance.Printf("%v", condition)
}

// Debug prints a debug message.
func (l *Logger) Debug(a ...interface{}) {
	instance.SetPrefix("[DEBUG] ")
	instance.Println(a...)
}

// Error prints an error message.
func (l *Logger) Error(a ...interface{}) {
	instance.SetPrefix("[ERROR] ")
	instance.Panicln(a...)
}

// Info prints an information message.
func (l *Logger) Info(a ...interface{}) {
	instance.SetPrefix("[INFO] ")
	instance.Println(a...)
}

// Warn prints a warning message.
func (l *Logger) Warn(a ...interface{}) {
	instance.SetPrefix("[WARNING] ")
	instance.Println(a...)
}