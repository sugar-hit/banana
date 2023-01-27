package web

import (
	"github.com/gin-gonic/gin"
)

// Reg :注册url路由与service业务流程之间的关系
func Reg() *gin.Engine {
	r := gin.Default()
	r.GET("/healthcheck", healthCheck)
	r.POST("/healthcheck", healthCheck)
	r.GET("/actuator/readiness", healthCheck)
	r.POST("/actuator/readiness", healthCheck)
	return r
}

func healthCheck(context *gin.Context) {
	context.JSON(200, gin.H{"message": "ok"})
}
