package stoppage

import "context"

type MockStopper struct {
	ErrToReturn, StopErrRoutine error
	CalledStop                  bool
}

func (m *MockStopper) Stop(ctx context.Context) error {
	m.CalledStop = true

	return m.StopErrRoutine
}

func (m MockStopper) StopError() error {
	return m.ErrToReturn
}
