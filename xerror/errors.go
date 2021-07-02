package xerror

var (
	Success         = New(0, "success", "")
	ErrUnknown      = New(999999, "unknown", "")
	ErrParamInvalid = New(100100, "param invalid", "")
)
