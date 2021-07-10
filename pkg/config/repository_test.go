package config

import (
	"context"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	repo := NewRepository()
	cases := []struct {
		key, value string
	}{
		{"ENV", "teste 1"},
		{"env", "teste 2"},
		{"foo_bar", "teste 3"},
		{"foo.bar", "teste 4"},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.key, func(t *testing.T) {
			repo.Add(tt.key, tt.value)

			resultedValue := repo.Get(tt.key)
			isSameValue := resultedValue != tt.value
			if isSameValue {
				t.Errorf(
					"got '%v' with key '%s', want '%s'",
					resultedValue,
					tt.key,
					tt.value,
				)
			}
		})
	}
}

func TestSource(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	repo := NewRepository()
	cases := []struct {
		key, value string
	}{
		{"env", "ENV"},
		{"foo_bar", "bar"},
		{"foo.bar", "foz"},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.key, func(t *testing.T) {
			repo.Source(ctx, MockSource{
				ReturnMap: map[string]string{
					tt.key: tt.value,
				},
			})

			resultedValue := repo.Get(tt.key)
			isSameValue := resultedValue != tt.value
			if isSameValue {
				t.Errorf(
					"expect source key %s with value %s, got value %s",
					tt.key,
					tt.value,
					resultedValue,
				)
			}
		})
	}
}
