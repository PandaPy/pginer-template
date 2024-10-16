package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code int         `json:"code"`           // 自定义的业务状态码
	Msg  string      `json:"msg"`            // 提示信息
	Data interface{} `json:"data,omitempty"` // 可选字段，包含具体数据
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{Code: 200, Msg: "Ok", Data: data})
}

func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, response{Code: errorCode, Msg: msg, Data: nil})
}

func FailByValidate(c *gin.Context, msg string) {
	Fail(c, HttpStatusInfo.ValidateError.ErrorCode, msg)
}

func FailByService(c *gin.Context, msg string) {
	Fail(c, HttpStatusInfo.ServiceError.ErrorCode, msg)
}

func FailByPanic(c *gin.Context, msg string) {
	Fail(c, HttpStatusInfo.PanicError.ErrorCode, msg)
}
