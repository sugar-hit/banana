package web

import (
	"github.com/gin-gonic/gin"
)

func Start(r *gin.Engine) {
	err := r.Run(":7001")
	if err != nil {
		return
	}
}
