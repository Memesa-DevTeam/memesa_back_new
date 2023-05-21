package middlewares

import (
	"github.com/gin-gonic/gin"
	"memesa_go_backend/app/service"
	"memesa_go_backend/pkg/api"
	"memesa_go_backend/pkg/authorizer"
	"memesa_go_backend/pkg/jwt"
	"net/http"
)

func Authorizer(cfg *jwt.JwtConfig, acfg *authorizer.AuthorizerConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get user token
		token := context.Request.Header.Get("Authorization")
		result := service.CheckTokenIsValid(token, context.Request.URL.String(), cfg, acfg)
		if !result {
			context.Next()
		}
		context.JSON(http.StatusInternalServerError, api.ReturnResponse(http.StatusInternalServerError, "Invalid Token Access"))
		context.Abort()
	}
}
