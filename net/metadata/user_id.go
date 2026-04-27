package metadata

import (
	"context"
)

const (
	ContextKeyUserID = "user_id"
)

func UserIdFromContext(ctx context.Context) (uid int64) {
	if v := ctx.Value(ContextKeyUserID); v != nil {
		if s, ok := v.(int64); ok {
			return s
		}
	}
	return 0
}
