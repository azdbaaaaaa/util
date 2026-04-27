package user

import (
	"github.com/azdbaaaaaa/util/proto/common"
)

const ContextKeyUser = "user"

type User struct {
	AppId    common.AppIdType `json:"app_id,omitempty"`
	UserId   int64            `json:"user_id,omitempty"`
	UserKey  string           `json:"user_key,omitempty"`
	Language string           `json:"language,omitempty"`
	Country  string           `json:"country,omitempty"`
}
