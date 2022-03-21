package metadata

import (
	"context"
)

const (
	ContextKeyTracing = "tracing"
)

func TracingFromContext(ctx context.Context) (s string) {
	if v := ctx.Value(ContextKeyTracing); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
