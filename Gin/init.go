package Gin

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var r *gin.Engine

func InitializeGinFramework(lc fx.Lifecycle) *gin.Engine {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			r = gin.Default()
			r.Use(cors.New(cors.Config{
				AllowAllOrigins:  true,
				AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "UPDATE"},
				AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
			}))
			// Middlewares Configuration

			// Router Groups configuration

			// API Configurations

			// Run program
			go func() {
				err := r.Run("0.0.0.0:8080")
				if err != nil {
					panic("Gin Error")
				}
			}()
			return nil
		},
	})
	return r
}
