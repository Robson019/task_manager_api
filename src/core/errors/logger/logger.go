package logger

import (
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"task_manager/src/core/errors"
	"time"
)

const (
	timeKey = "time"
	fileKey = "file"
	lineKey = "line"
)

type Logger interface {
	Log(err errors.Error)
	LogWithMessage(logLevel int, message string)
}

type TaskManagerLogger struct {
	zeroLogger zerolog.Logger
}

func (instance TaskManagerLogger) Log(err errors.Error) {
	instance.getEventLog(err.LogLevel()).
		Time(timeKey, time.Now()).
		Str(fileKey, err.File()).
		Int(lineKey, err.Line()).
		Msg(err.Error())
}

func (instance TaskManagerLogger) LogWithMessage(logLevel int, message string) {
	_, file, line, _ := runtime.Caller(1)
	instance.getEventLog(logLevel).
		Time(timeKey, time.Now()).
		Str(fileKey, file).
		Int(lineKey, line).
		Msg(message)
}

func (instance TaskManagerLogger) getEventLog(level int) *zerolog.Event {
	switch level {
	case errors.DebugLevel:
		return instance.zeroLogger.Debug()
	case errors.WarnLevel:
		return instance.zeroLogger.Warn()
	case errors.ErrorLevel:
		return instance.zeroLogger.Error()
	case errors.FatalLevel:
		return instance.zeroLogger.Fatal()
	default:
		return instance.zeroLogger.Info()
	}
}

func New() *TaskManagerLogger {
	logger := zerolog.New(os.Stdout)
	return &TaskManagerLogger{zeroLogger: logger}
}
