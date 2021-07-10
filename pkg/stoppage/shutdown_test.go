package stoppage

import (
	"context"
	"errors"
	"testing"

	"fake-hash.ports-adapters/pkg/log"
)

func TestGraceful(t *testing.T) {
	t.Parallel()

	stopErr := errors.New("error on stop routine")
	errToReturn := errors.New("stop error reason")

	logger := &log.MockLogger{}
	ctx := context.Background()

	cases := []struct {
		name              string
		stopper           *MockStopper
		expectedErrReason string
	}{
		{"Graceful stop with no error", &MockStopper{StopErrRoutine: nil}, ""},
		{"Graceful stop with error from Stop method", &MockStopper{StopErrRoutine: stopErr}, stopErr.Error()},
		{"Graceful stop with error given a defined reason", &MockStopper{StopErrRoutine: stopErr, ErrToReturn: errToReturn}, errToReturn.Error()},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			shutdown := New(logger)
			shutdown.Add(tt.stopper)
			shutdown.Graceful(ctx)

			if !tt.stopper.CalledStop {
				t.Errorf("shutdown should have called the Stop method")
			}

			if tt.expectedErrReason != logger.ErrorMessage {
				t.Errorf("expect error reason '%s'; got '%s'", tt.expectedErrReason, logger.ErrorMessage)
			}
		})
	}
}
