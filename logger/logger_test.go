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

func TestLogger_CheckLogPrefix(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()
	log := "test log"
	logger.Log(log)
	ass.Equal("[LOG] ", logger.Prefix())
}

func TestLogger_CheckAssertPrefix(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()

	logger.Assert(true)
	ass.Equal("[ASSERT] ", logger.Prefix())
}

func TestLogger_CheckDebugPrefix(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()

	logger.Debug("Test")
	ass.Equal("[DEBUG] ", logger.Prefix())
}

func TestLogger_CheckErrorPrefix(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()

	ass.Panics(func() {logger.Error("Log error")})
	ass.Equal("[ERROR] ", logger.Prefix())
}

func TestLogger_CheckPanicWithMessage(t *testing.T) {
	ass    := assert.New(t)
	logger := GetInstance()

	log := "error log"
	ass.PanicsWithValue(log, func() {logger.Error(log)})

}

func TestLogger_CheckInfoPrefix(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()

	logger.Info("Info test")
	ass.Equal("[INFO] ", logger.Prefix())
}

func TestLogger_CheckWarnPrefix(t *testing.T) {
	ass := assert.New(t)
	logger := GetInstance()

	logger.Warn("Info warn")
	ass.Equal("[WARNING] ", logger.Prefix())
}

