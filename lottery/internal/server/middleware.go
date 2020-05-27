package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func accessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		//r:=c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		c.Next()
	}
}

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !svc.ReqLimiter.Allow() {
			msg := "系统繁忙，限流中"
			svc.GetSysLogger().Println(msg)
			c.JSON(http.StatusOK, msg)
			c.Abort()
		}
		c.Next()
	}
}