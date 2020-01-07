package logger

import (
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

func (l *Logger) Error(message string) {
	instance.SetPrefix("[ERROR] ")
	instance.Println(message)
}

func (l *Logger) Info(message string) {
	instance.SetPrefix("[INFO] ")
	instance.Println(message)
}

func (l *Logger) Warn(message string) {
	instance.SetPrefix("[WARNING] ")
	instance.Println(message)
}