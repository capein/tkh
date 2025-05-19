// Package logger provides a logger that logs messages with context
package logger

import (
	"arc/logger/contextLogger"
	"context"
	"runtime"
)

type ILogger interface {
	Println(file string, line int, args ...interface{})
	Debug(file string, line int, args ...interface{})
	Error(file string, line int, args ...interface{})
	Info(file string, line int, args ...interface{})
	Warn(file string, line int, args ...interface{})
	RegisterContextHandler(ctx context.Context, name string, handler func(ctx context.Context) map[string]interface{}) error
}

var l ILogger = &contextLogger.Logger{}

func Println(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Println(file, line, args...)
}

func Debug(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Debug(file, line, args...)
}

func Error(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Error(file, line, args...)
}

func Warn(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Warn(file, line, args...)
}

func Info(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.Info(file, line, args...)
}

func RegisterContextHandler(ctx context.Context, name string, handler func(ctx context.Context) map[string]interface{}) error {
	return l.RegisterContextHandler(ctx, name, handler)
}
