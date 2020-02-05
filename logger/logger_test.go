package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestLogger_New(t *testing.T) {
	ass := assert.New(t)
	logger := New()
	ass.IsType(&Logger{}, logger)
}

func TestLogger_GetInstance(t *testing.T) {
	ass := assert.New(t)
	instanceLogger := GetInstance()
	ass.IsType(&Logger{}, instanceLogger)
}

func TestLogger_GetInstance_SamePointer(t *testing.T) {
	ass := assert.New(t)
	instanceLogger1 := GetInstance()
	instanceLogger2 := GetInstance()
	ass.Equal(instanceLogger1, instanceLogger2)
}

func TestLogger_LogPrepend(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()
	log := "test log"
	logger.Log(log)
	ass.Equal("[LOG] ", logger.Prefix())
}