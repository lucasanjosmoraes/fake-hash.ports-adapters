package log

import (
	"context"
	"fmt"
)

type MockLogger struct {
	DebugMessage, ErrorMessage, InfoMessage, PanicMessage string
}

func (m *MockLogger) Debug(ctx context.Context, args ...interface{}) {
	m.DebugMessage = fmt.Sprintf("%v", args...)
}

func (m *MockLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	m.DebugMessage = fmt.Sprintf(format, args...)
}

func (m *MockLogger) Error(ctx context.Context, args ...interface{}) {
	m.ErrorMessage = fmt.Sprintf("%v", args...)
}

func (m *MockLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	m.ErrorMessage = fmt.Sprintf(format, args...)
}

func (m *MockLogger) Info(ctx context.Context, args ...interface{}) {
	m.InfoMessage = fmt.Sprintf("%v", args...)
}

func (m *MockLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	m.InfoMessage = fmt.Sprintf(format, args...)
}

func (m *MockLogger) Panic(ctx context.Context, args ...interface{}) {
	m.PanicMessage = fmt.Sprintf("%v", args...)
}

func (m *MockLogger) Panicf(ctx context.Context, format string, args ...interface{}) {
	m.PanicMessage = fmt.Sprintf(format, args...)
}
