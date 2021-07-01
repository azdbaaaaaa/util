package xcode

var (
	ErrCodeSuccess = NewCodeError(0, "成功")
	ErrCodeUnknown = NewCodeError(999999, "未知错误")
)
