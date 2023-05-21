package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	authorizer2 "memesa_go_backend/app/middlewares/authorizer"
	"memesa_go_backend/pkg/jwt"
	"memesa_go_backend/pkg/server"
)

func Provide() fx.Option {
	return fx.Provide(InitRouters)
}

func InitRouters(authorizer *jwt.Authorizer) server.InitRouter {
	return func(r *gin.Engine) {
		// Middlewares
		r.Use(authorizer2.CheckTokenIsValid(authorizer))

		// Test router
		r.GET("/greetings", func(c *gin.Context) {
			c.JSON(200, "Hello!")
		})
	}
}
