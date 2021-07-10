package errorhandler

import (
	"context"

	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/log"
)

// Logger define o que um erro precisa implementar para logar suas informações.
type Logger interface {
	error
	Log(ctx context.Context, l log.Logger)
}
