package errorhandler

import (
	"context"

	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/http"
	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/log"
)

const DefaultErrorMessage = "unkwnow internal error"

// HTTP fornece uma implementação padrão do HTTPHandler.
type HTTP struct {
	Logger log.Logger
}

// New instancia HTTP, mas também nos ajuda a validar se ele implementa
// HTTPHandler corretamente.
func New(l log.Logger) HTTPHandler {
	return HTTP{
		Logger: l,
	}
}

// Handle valida se o erro dado implementa Logger e HTTPResponder para então
// chamar seus métodos.
func (eh HTTP) Handle(ctx context.Context, r http.Response, err error) {
	eh.log(ctx, err)
	eh.response(r, err)
}

func (eh HTTP) log(ctx context.Context, err error) {
	if err == nil {
		return
	}

	logErr, ok := err.(Logger)
	if ok {
		logErr.Log(ctx, eh.Logger)
		return
	}

	eh.Logger.Error(ctx, err)
}

func (eh HTTP) response(r http.Response, err error) {
	appErr, ok := err.(HTTPResponder)
	if ok {
		_ = r.Write(appErr.Status(), appErr.Response())
		return
	}

	_ = r.WriteInternalError([]byte(DefaultErrorMessage))
}
