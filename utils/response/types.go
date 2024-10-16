package response

type ResponseStatus struct {
	ErrorCode int
	ErrorMsg  string
}

type HttpStatus struct {
	Success       ResponseStatus
	ValidateError ResponseStatus
	ServiceError  ResponseStatus
	PanicError    ResponseStatus
}

var HttpStatusInfo = HttpStatus{
	Success:       ResponseStatus{200, "操作成功"},
	ValidateError: ResponseStatus{4001, "请求参数错误"},
	ServiceError:  ResponseStatus{4002, "业务错误"},
	PanicError:    ResponseStatus{4004, "系统异常，请稍后重试"},
}
