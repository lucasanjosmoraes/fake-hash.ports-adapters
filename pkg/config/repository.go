package config

import "context"

// Repository é uma entidade que nos ajuda a gerenciar variáveis retornadas de
// um Sourcer.
type Repository struct {
	values map[string]string
}

func NewRepository() *Repository {
	return &Repository{
		values: make(map[string]string),
	}
}

// Source carregará todas as variáveis retornadas de um Sourcer em um Repository.
func (r *Repository) Source(ctx context.Context, getter Sourcer) {
	for k, v := range getter.Load(ctx) {
		r.Add(k, v)
	}
}

// Add adicionará a chave e valor dados ao Repository.
func (r *Repository) Add(k string, v string) {
	r.values[k] = v
}

// Get procurará pela chave dada no map de valores do Repository.
func (r Repository) Get(key string) string {
	return r.values[key]
}
