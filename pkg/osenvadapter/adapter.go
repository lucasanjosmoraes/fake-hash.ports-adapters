package osenvadapter

import (
	"context"
	"os"
	"strings"

	"fake-hash.ports-adapters/pkg/config"
)

// Adapter implementa config.Sourcer, usando o pacote "os" como fonte de dados.
type Adapter struct {
	prefix string
}

// New instancia um novo Adapter, mas também nos ajuda a validar se ele implemeta
// config.Sourcer corretamente.
func New(prefix string) config.Sourcer {
	return Adapter{
		prefix: prefix,
	}
}

// Load irá buscar por todas as variáveis de ambiente do sistema e as carregará
// em um novo map.
func (a Adapter) Load(ctx context.Context) map[string]string {
	values := make(map[string]string)

	for _, env := range os.Environ() {
		sepEnv := strings.SplitN(env, "=", 2)
		k := sepEnv[0]
		v := sepEnv[1]

		if strings.HasPrefix(k, a.prefix) {
			k = strings.Replace(k, a.prefix, "", 1)
			values[k] = v
		}
	}

	return values
}
