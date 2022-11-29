package router

import (
	"github.com/gin-gonic/gin"
	"myself/controller"
	"net/http"
)

func Setup() (r *gin.Engine, err error) {
	r = gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong!")
	})
	r.POST("/signup", controller.SignUpHandler)
	r.Group("api/talk")
	return r, nil
}
