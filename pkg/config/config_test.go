package config

import "context"

type MockSource struct {
	ReturnMap map[string]string
}

func (m MockSource) Load(ctx context.Context) map[string]string {
	return m.ReturnMap
}
