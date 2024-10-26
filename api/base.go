package api

const (
	JsonContentType   = "application/json"
	MaxRequestElemNum = 50

	ErrCodeSuccess      = 0
	ErrCodeUnknownError = 1
	ErrCodeInvalidParam = 6
	ErrCodeLessAuth     = 109 // 权限不足
	ErrCodeAuthFailed   = 110 // 鉴权失败
)

type Request interface {
}

type BaseRequest struct {
	IatSec int64 `json:"iat"`
	ExpSec int64 `json:"exp"`
}

func (req *BaseRequest) SetIat(iat int64) {
	req.IatSec = iat
}

func (req *BaseRequest) Iat() int64 {
	return req.IatSec
}

func (req *BaseRequest) SetExp(exp int64) {
	req.ExpSec = exp
}

func (req *BaseRequest) Exp() int64 {
	return req.ExpSec
}

type Response[DataT any] struct {
	ErrCode  int32  `json:"bizcode"`
	ErrorMsg string `json:"error"`
	Data     DataT  `json:"data"`
}
