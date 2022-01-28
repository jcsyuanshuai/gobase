package output

import (
	"context"
	"net/http"
)

type Enum interface {
	WithError(err error) Enum
	GetError() error

	WithContext(ctx context.Context) Enum
	Context() context.Context

	WithCode(code int) Enum
	GetCode() int

	WithMsg(desc string) Enum
	GetMsg() string

	WithHttpCode(code int) Enum
	GetHttpCode() int
}

type enum struct {
	Code     int             // 业务码
	Msg      string          // 美化描述
	ctx      context.Context //上下文
	err      error           //发生的错误
	HttpCode int
}

func (e enum) WithError(err error) Enum {
	e.err = err
	return e
}

func (e enum) GetError() error {
	return e.err
}

func (e enum) WithContext(ctx context.Context) Enum {
	e.ctx = ctx
	return e
}

func (e enum) Context() context.Context {
	if e.ctx != nil {
		return e.ctx
	}
	return context.Background()
}

func (e enum) WithCode(code int) Enum {
	e.Code = code
	return e
}

func (e enum) GetCode() int {
	return e.Code
}

func (e enum) WithMsg(desc string) Enum {
	e.Msg = desc
	return e
}

func (e enum) GetMsg() string {
	return e.Msg
}

func (e enum) WithHttpCode(code int) Enum {
	e.HttpCode = code
	return e
}

func (e enum) GetHttpCode() int {
	if e.HttpCode == 0 {
		return 200
	}
	return e.HttpCode
}

var (
	Success Enum = enum{
		Code:     10000,
		HttpCode: http.StatusOK,
		Msg:      "success",
	}
)
