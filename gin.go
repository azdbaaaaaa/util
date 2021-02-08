package util

type CommonResponse struct {
	Code    Code                   `json:"code"`
	Msg     string                 `json:"msg"`
	Success bool                   `json:"success"`
	Result  interface{}            `json:"result"`
	Tid     string                 `json:"tid"`
	Ext     map[string]interface{} `json:"ext,omitempty"`
}

//
//type Response struct {
//	Errno int         `json:"errno"`
//	Msg   string      `json:"msg"`
//	Data  interface{} `json:"data,omitempty"`
//}
//
//func Send(c *gin.Context, code int, msg string, data interface{}) {
//	c.JSON(http.StatusOK, Response{
//		Errno: code,
//		Msg:   msg,
//		Data:  data,
//	})
//}
//
//func Failure(c *gin.Context, err error) {
//	c.JSON(http.StatusOK, Response{
//		Errno: -1,
//		Msg:   "error",
//		Data:  err.Error(),
//	})
//}
//
//func Success(c *gin.Context, data interface{}) {
//	c.JSON(http.StatusOK, Response{
//		Errno: 0,
//		Msg:   "success",
//		Data:  data,
//	})
//}
