package toolkit

import (
	"context"

	"github.com/lucasanjosmoraes/fake-hash.ports-adapters/pkg/http"
)

const (
	CorrelationKey       = "correlationID"
	CorrelationHeaderKey = "x-correlation-id"
)

// LoadCorrelationID procura pelo header relacionado ao correlation ID na
// requisição HTTP dada e o carrega no contexto dado.
func LoadCorrelationID(ctx context.Context, req http.Request) context.Context {
	prev := ExtractCorrelationID(ctx)
	correlationID := req.Header(CorrelationHeaderKey)

	if correlationID == "" && prev != "" {
		correlationID = prev
	}

	return context.WithValue(ctx, CorrelationKey, correlationID)
}

// ExtractCorrelationID irá procurar pelo correlation ID no contexto dado.
func ExtractCorrelationID(ctx context.Context) string {
	v := ctx.Value(CorrelationKey)
	if v == nil {
		return ""
	}

	c, ok := v.(string)
	if ok {
		return c
	}

	return ""
}
