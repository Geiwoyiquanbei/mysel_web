package middleware

import (
	"github.com/gin-gonic/gin"
	"myself/pkg/JWt"
)

// 中间件 jwt
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 假设 token 的是在 url 中存放的
		atoken := c.DefaultQuery("atoken", "")
		ftoken := c.DefaultQuery("ftoken", "")
		if atoken == "" || ftoken == "" {
			c.JSON(200, gin.H{
				"msg":  "没有携带 atoken 或者 ftoken",
				"code": 400,
			})
			c.Abort()
			return
		}
		parasToken, err := JWt.ParasToken2(atoken) // 解析 access_token
		if err == nil {                            // 当前的 access_token 格式对，没有过期
			c.JSON(200, gin.H{
				"msg":  "atoken 和 ftoken 没有过期",
				"data": parasToken,
				"code": 400,
			})
			c.Next()
			return
		}
		atoken, rToken, err := JWt.RefreshToken(atoken, ftoken)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "您需要重新登录",
				"code": 400,
			})
			c.Abort()
			return
		} else {
			c.JSON(200, gin.H{
				"msg":    "atoken 和 ftoken 没有过期",
				"atoken": atoken,
				"rToken": rToken,
				"code":   400,
			})
			c.Next()
			return
		}
	}
}
