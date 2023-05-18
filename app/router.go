package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"memesa_go_backend/pkg/server"
)

func Provide() fx.Option {
	return fx.Provide(InitRouters)
}

func InitRouters() server.InitRouter {
	return func(r *gin.Engine) {
		// Test router
		r.GET("/greetings", func(c *gin.Context) {
			c.JSON(200, "Hello!")
		})
	}
}
