package log

import (
	"context"
)

type Level int

const (
	Debug Level = iota + 1
	Error
	Info
	Panic
)

// Logger define tudo que um Adapter necessita fornecer para um consumidor de log.
type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Panic(ctx context.Context, args ...interface{})
	Panicf(ctx context.Context, format string, args ...interface{})
}
