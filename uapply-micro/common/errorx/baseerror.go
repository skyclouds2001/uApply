package errorx

const (
	Success ErrorInt = 1000 + iota
	Fail
	ParamInvalid // 参数错误
	SystemBusy   // 系统繁忙
	NotFound     // 数据库无数据
	Exist        // 索引冲突
	PasswordInvalid
	NotAuth // 没有权限
)

type ErrorInt int

type CodeError struct {
	Code ErrorInt `json:"code"`
	Msg  string   `json:"msg"`
}

type CodeErrorResponse struct {
	Code ErrorInt `json:"code"`
	Msg  string   `json:"msg"`
}

func NewCodeError(code ErrorInt, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewError(msg string, code ErrorInt) error {
	return NewCodeError(code, msg)
}

func NewSysError() error {
	return NewCodeError(SystemBusy, "系统繁忙")
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
