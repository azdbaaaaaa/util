package metadata

import (
	"context"
)

const (
	ContextKeyReqID = "req_id"
)

func ReqIdFromContext(ctx context.Context) (s string) {
	if v := ctx.Value(ContextKeyInParam); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
