package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type Server struct {
	router *gin.Engine
	server *http.Server
}

// Start the server
func (s *Server) Start() error {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Println("Http Server Error: ", err)
		}
		fmt.Println("Server is running at ", s.server.Addr)
	}()
	return nil
}

// Stop the server
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// try to shutdown the server
	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Println("Unable to shutdown the server...")
	}
	return nil
}

// Cors
// Cors middleware to use if enabled
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CORS is ALLOWED on this server.")
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func NewServer(router *gin.Engine, config *ServerConfig) *Server {
	return &Server{
		router: router,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Address, config.Port),
			Handler: router,
		},
	}
}

type InitRouter func(r *gin.Engine)

func NewRouter(config *ServerConfig, init InitRouter) *gin.Engine {
	router := gin.New()
	// Set up Cors Policies
	if config.AllowCORS {
		router.Use(Cors())
	}
	// Set up recovery
	router.Use(gin.Recovery())
	// Initialize Routers
	init(router)
	return router
}

// lifecycle
func lc(lifecycle fx.Lifecycle, server *Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return server.Start()
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop()
		},
	})
}

func Provide() fx.Option {
	return fx.Options(fx.Provide(NewServerConfig, NewServer, NewRouter), fx.Invoke(lc))
}
