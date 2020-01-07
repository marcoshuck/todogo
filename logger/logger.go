package logger

import (
	"fmt"
	"github.com/marcoshuck/todogo/errors"
	"log"
	"os"
	"sync"
)

var instance *Logger
var once sync.Once

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[LOG]", log.LstdFlags),
	}
}

func instantiate() {
	instance = New()
}

func GetInstance() *Logger {
	once.Do(instantiate)
	return instance
}

func (l *Logger) Log(message string) {
	instance.SetPrefix("[LOG] ")
	instance.Println(message)
}

func (l *Logger) Assert(condition bool) {
	instance.SetPrefix("[ASSERT] ")
	instance.Printf("%v", condition)
}

func (l *Logger) Debug(message string) {
	instance.SetPrefix("[DEBUG] ")
	instance.Println(message)
}

func (l *Logger) Error(err errors.Error) {
	instance.SetPrefix("[ERROR] ")
	instance.Println(fmt.Sprintf("[Error] Status: %d | Code: %d | Message: %s", err.Status, err.Code, err.Message))
	instance.Println(fmt.Sprintf("Stack trace: %v", err.Base))
}

func (l *Logger) Info(message string) {
	instance.SetPrefix("[INFO] ")
	instance.Println(message)
}

func (l *Logger) Warn(message string) {
	instance.SetPrefix("[WARNING] ")
	instance.Println(message)
}