package config

import "context"

// Sourcer pode ser usada para carregar variáveis de ambiente de uma fonte definida.
type Sourcer interface {
	Load(ctx context.Context) map[string]string
}
