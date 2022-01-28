package output

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Result struct {
	Code int
	Msg  string
	Data interface{}
	time time.Time
}

func Json(ctx *gin.Context, enum Enum, data interface{}) {
	ctx.JSON(enum.GetHttpCode(), &Result{
		Code: enum.GetCode(),
		Msg:  enum.GetMsg(),
		Data: data,
		time: time.Now(),
	})
}
