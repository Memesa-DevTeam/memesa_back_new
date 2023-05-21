package middlewares

import (
	"github.com/gin-gonic/gin"
)

type InitMiddlewares func(r *gin.Engine)

func InitMiddleware() InitMiddlewares {
	return func(r *gin.Engine) {
		// Insert all middlewares that affect globally here

	}
}
