package helpers

import (
	"github.com/gin-gonic/gin"
)

func ResponseJSON(ctx *gin.Context, code int, message string, payload interface{}) {
	ctx.JSON(code, gin.H{
		"code_status": code,
		"message":     message,
		"payload":     payload,
	})
}
