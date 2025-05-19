package basic

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Logger struct {
}

func (l Logger) Error(file string, line int, args ...interface{}) {
	data := make([]interface{}, 0, len(args)+2)
	data = append(data, fmt.Sprintf("%s:%d", file, line))
	data = append(data, args...)
	log.Println(data...)
}

func (l Logger) Info(file string, line int, args ...interface{}) {
	data := make([]interface{}, 0, len(args)+2)
	data = append(data, fmt.Sprintf("%s:%d", file, line))
	data = append(data, args...)
	log.Println(data...)
}

func (l Logger) Warn(file string, line int, args ...interface{}) {
	data := make([]interface{}, 0, len(args)+2)
	data = append(data, fmt.Sprintf("%s:%d", file, line))
	data = append(data, args...)
	log.Println(data...)
}

func (l Logger) RegisterContextHandler(ctx context.Context, name string, handler func(ctx context.Context) map[string]interface{}) error {
	return nil
}

func (l Logger) Println(file string, line int, args ...interface{}) {
	data := make([]interface{}, 0, len(args)+2)
	data = append(data, fmt.Sprintf("%s:%d", file, line))
	data = append(data, args...)
	log.Println(data...)
}

func (l Logger) Debug(file string, line int, args ...interface{}) {
	if os.Getenv("debug") == "" {
		return
	}
	data := make([]interface{}, 0, len(args)+2)
	data = append(data, fmt.Sprintf("%s:%d", file, line))
	data = append(data, args...)
	log.Println(data...)
}
