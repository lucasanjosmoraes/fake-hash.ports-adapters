package errorhandler

import (
	"context"
	"errors"
	nethttp "net/http"
	"testing"

	"fake-hash.ports-adapters/pkg/http"
	"fake-hash.ports-adapters/pkg/log"
)

func TestHandle(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	rawMessage := "raw error"
	debugMessage := "debug message"
	infoMessage := "info message"
	errorMessage := "error message"

	rawErr := errors.New(rawMessage)
	debugErr := LoggableError{LogLevel: log.Debug, LogMessage: debugMessage, Err: debugMessage}
	infoErr := LoggableError{LogLevel: log.Info, LogMessage: infoMessage, Err: infoMessage}
	errorErr := LoggableError{LogLevel: log.Error, LogMessage: errorMessage, Err: errorMessage}

	okBody := []byte("ok response")
	createdBody := []byte("created body")
	badRequestBody := []byte("bad request body")

	okErr := HTTPError{StatusCode: nethttp.StatusOK, Body: okBody}
	createdErr := HTTPError{StatusCode: nethttp.StatusCreated, Body: createdBody}
	badRequestErr := HTTPError{StatusCode: nethttp.StatusBadRequest, Body: badRequestBody}

	logger := &log.MockLogger{}
	handler := New(logger)
	response := &MockResponse{}

	defaultBody := []byte(DefaultErrorMessage)

	cases := []struct {
		name           string
		logger         log.Logger
		handler        HTTPHandler
		error          error
		expectedLevel  log.Level
		response       http.Response
		expectedStatus int
		expectedBody   []byte
	}{
		{"Log and response raw error", logger, handler, rawErr, log.Error, response, nethttp.StatusInternalServerError, defaultBody},
		{"Log error at Debug Level", logger, handler, debugErr, log.Debug, response, nethttp.StatusInternalServerError, defaultBody},
		{"Log error at Info Level", logger, handler, infoErr, log.Info, response, nethttp.StatusInternalServerError, defaultBody},
		{"Log error at Error Level", logger, handler, errorErr, log.Error, response, nethttp.StatusInternalServerError, defaultBody},
		{"Response error with ok status", logger, handler, okErr, log.Error, response, nethttp.StatusOK, okBody},
		{"Response error with created status", logger, handler, createdErr, log.Error, response, nethttp.StatusCreated, createdBody},
		{"Response error with bad request status", logger, handler, badRequestErr, log.Error, response, nethttp.StatusBadRequest, badRequestBody},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.handler.Handle(ctx, tt.response, tt.error)

			switch tt.expectedLevel {
			case log.Debug:
				if logger.DebugMessage != tt.error.Error() {
					t.Errorf("expect '%s' message at Debug level; resulted as %+v", tt.error.Error(), logger)
				}
			case log.Info:
				if logger.InfoMessage != tt.error.Error() {
					t.Errorf("expect '%s' message at Info level; resulted as %+v", tt.error.Error(), logger)
				}
			case log.Error:
				if logger.ErrorMessage != tt.error.Error() {
					t.Errorf("expect '%s' message at Error level; resulted as %+v", tt.error.Error(), logger)
				}
			case log.Panic:
				if logger.PanicMessage != tt.error.Error() {
					t.Errorf("expect '%s' message at Panic level; resulted as %+v", tt.error.Error(), logger)
				}
			}

			if response.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, response.StatusCode)
			}

			resultedBody := string(response.Body)
			expectedBody := string(tt.expectedBody)
			if resultedBody != expectedBody {
				t.Errorf("expected body %s; got %s", expectedBody, resultedBody)
			}
		})
	}
}
