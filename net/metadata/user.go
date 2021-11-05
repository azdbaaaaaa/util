package metadata

import (
	"context"
	"github.com/azdbaaaaaa/util/xutil/xerror"
)

const (
	ContextKeyUser = "user"
)

type User struct {
	UserId int64 `json:"user_id,omitempty"`
}

func UserFromContext(ctx context.Context) (u User, err error) {
	if v := ctx.Value(ContextKeyUser); v != nil {
		if u, ok := v.(User); ok {
			return u, nil
		}
	}
	return u, xerror.ErrNoUserError
}
