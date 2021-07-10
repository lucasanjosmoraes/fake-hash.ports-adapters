package toolkit

import (
	"context"
	"testing"
)

func TestLoadCorrelationID(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name              string
		context           context.Context
		correlationHeader map[string]string
		expectCorrelation string
	}{
		{
			"Empty context",
			context.Background(),
			map[string]string{},
			"",
		},
		{
			"Context with no correlation ID",
			context.Background(),
			map[string]string{"foo": "123"},
			""},
		{
			"Successful case",
			context.Background(),
			map[string]string{CorrelationHeaderKey: "123"},
			"123",
		},
		{
			"Successful case with previous correlation ID only",
			context.WithValue(context.Background(), CorrelationKey, "123"),
			map[string]string{},
			"123",
		},
		{
			"Successful case with previous correlation ID",
			context.WithValue(context.Background(), CorrelationKey, "123"),
			map[string]string{CorrelationHeaderKey: "456"},
			"456",
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := MockHTTPRequest{
				Headers: tt.correlationHeader,
			}
			ctx := LoadCorrelationID(tt.context, req)

			var correlation string
			val := ctx.Value(CorrelationKey)
			if val != nil {
				correlation = val.(string)
			}

			if correlation != tt.expectCorrelation {
				t.Errorf("expect correlationID %s got %s", tt.expectCorrelation, correlation)
			}
		})
	}
}

func TestExtractCorrelationID(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name              string
		context           context.Context
		expectCorrelation string
	}{
		{"Empty correlation ID on context.Background", context.Background(), ""},
		{"Empty correlation ID on context.TODO", context.TODO(), ""},
		{"Successful case", context.WithValue(context.Background(), CorrelationKey, "123"), "123"},
		{"Successful case even when correlation ID is empty", context.WithValue(context.Background(), CorrelationKey, ""), ""},
		{"Empty correlation ID when correlation ID has a invalid format", context.WithValue(context.Background(), CorrelationKey, 123), ""},
	}

	for _, tt := range cases {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			resultedCorrelationID := ExtractCorrelationID(tt.context)

			if resultedCorrelationID != tt.expectCorrelation {
				t.Errorf("expect correlationID %s; got %s", tt.expectCorrelation, resultedCorrelationID)
			}
		})
	}
}
