package errorhandler

import (
	"context"
	"net/http"

	"fake-hash.ports-adapters/pkg/log"
)

type LoggableError struct {
	Err        string
	LogMessage string
	LogLevel   log.Level
}

func (e LoggableError) Error() string {
	return e.Err
}

func (e LoggableError) Log(ctx context.Context, l log.Logger) {
	switch e.LogLevel {
	case log.Debug:
		l.Debug(ctx, e.LogMessage)
	case log.Info:
		l.Info(ctx, e.LogMessage)
	case log.Error:
		l.Error(ctx, e.LogMessage)
	case log.Panic:
		l.Panic(ctx, e.LogMessage)
	}
}

type HTTPError struct {
	Err        string
	StatusCode int
	Body       []byte
}

func (e HTTPError) Error() string {
	return e.Err
}

func (e HTTPError) Status() int {
	return e.StatusCode
}

func (e HTTPError) Response() []byte {
	return e.Body
}

type MockResponse struct {
	StatusCode   int
	Body         []byte
	JSONResponse bool
}

func (m *MockResponse) Write(statusCode int, body []byte) error {
	m.StatusCode = statusCode
	m.Body = body

	return nil
}

func (m *MockResponse) WriteCreated(body []byte) error {
	m.StatusCode = http.StatusCreated
	m.Body = body

	return nil
}

func (m *MockResponse) WriteBadRequest(body []byte) error {
	m.StatusCode = http.StatusBadRequest
	m.Body = body

	return nil
}

func (m *MockResponse) WriteInternalError(body []byte) error {
	m.StatusCode = http.StatusInternalServerError
	m.Body = body

	return nil
}

func (m *MockResponse) JsonResponse() {
	m.JSONResponse = true
}
