package stoppage

import "context"

// Stopper define como Adapters devem gerenciar suas rotinas de shutdown.
type Stopper interface {
	Stop(ctx context.Context) error
	StopError() error
}
