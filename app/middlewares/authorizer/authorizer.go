package authorizer

import (
	"github.com/gin-gonic/gin"
	"memesa_go_backend/pkg/api"
	"memesa_go_backend/pkg/jwt"
	"net/http"
	"regexp"
)

func CheckTokenIsValid(a *jwt.Authorizer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// iterate the whole list to find the pattern
		targetURL := ctx.Request.URL
		token := ctx.Request.Header.Get("Authorization")
		for i := 0; i < len(a.Exclude); i++ {
			if targetURL.String() == a.Exclude[i] {
				ctx.Next()
			}
			// Generate pattern
			pattern := regexp.MustCompile(a.Exclude[i] + `.*`)
			if pattern.MatchString(targetURL.String()) {
				ctx.Next()
			}
		}
		// Check token
		if !a.CheckIsValid(token) {
			ctx.JSON(http.StatusInternalServerError, api.ReturnResponse(http.StatusInternalServerError, "Invalid token"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
