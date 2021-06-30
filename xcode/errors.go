package xcode

var (
	Success      = NewCode(0, "成功")
	ErrorUnknown = NewCode(999999, "未知错误")
)

type ErrorCode struct {
	// 错误码 ex:999999
	Code int `json:"code"`
	// 错误消息
	Msg string `json:"msg"`
	// 数据
	Data string `json:"data"`

	// error debug信息
	Error string `json:"error"`
	// request id
	Rid string `json:"rid"`
}

//func ()  {
//
//}