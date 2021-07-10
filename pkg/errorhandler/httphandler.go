package errorhandler

import (
	"context"

	"fake-hash.ports-adapters/pkg/http"
)

// HTTPHandler define o que é necessário para gerenciar erros retornados do http.Response.
type HTTPHandler interface {
	Handle(ctx context.Context, r http.Response, err error)
}
