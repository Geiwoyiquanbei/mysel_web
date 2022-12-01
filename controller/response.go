package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseWithError(c *gin.Context, code ResCode) {
	var rd = &ResData{
		Code: code,
		Msg:  Codemap[code],
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseWithSuccess(c *gin.Context, code ResCode) {
	var rd = &ResData{
		Code: code,
		Msg:  Codemap[code],
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseWithMsg(c *gin.Context, code ResCode, data interface{}) {
	var rd = &ResData{
		Code: code,
		Msg:  Codemap[code],
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
