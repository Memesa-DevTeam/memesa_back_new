package service

import (
	"fmt"
	"memesa_go_backend/pkg/authorizer"
	"memesa_go_backend/pkg/jwt"
	"regexp"
)

func CheckTokenIsValid(token string, targetURL string, config *jwt.JwtConfig, authConfig *authorizer.AuthorizerConfig) bool {
	// iterate to see whether the request is in the exclude paths
	for i := 0; i < len(authConfig.ExcludePaths); i++ {
		// check patterns
		pattern, _ := regexp.Compile(authConfig.ExcludePaths[i] + `.*`)
		if pattern.MatchString(targetURL) {
			// skip checking
			fmt.Println("Excluded Path Checked. Passing...")
			return true
		}
		if targetURL == authConfig.ExcludePaths[i] {
			fmt.Println("Excluded Path Checked. Passing...")
			return true
		}
	}
	if !jwt.CheckIsValid(token, config) {
		return false
	}
	return true
}
