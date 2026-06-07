// internal/middleware/auth.go
package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"Backend-Warehouse/domain/repository"
)

type AuthMiddleware struct {
	jwtService repository.JWTService
}

func NewAuthMiddleware(jwt repository.JWTService) *AuthMiddleware {
	return &AuthMiddleware{jwtService: jwt}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, ok := extractBearerToken(c)
		if !ok {
			abortWithError(c, http.StatusUnauthorized, "missing or invalid authorization header", "TOKEN_MISSING")
			return
		}

		claims, err := m.jwtService.ValidateAccessToken(c.Request.Context(), accessToken)
		if err != nil {
			switch {
                case errors.Is(err, repository.ErrTokenExpired):
                    abortWithError(c, http.StatusUnauthorized, "access token expired", "TOKEN_EXPIRED")
                case errors.Is(err, repository.ErrTokenRevoked):
                    abortWithError(c, http.StatusUnauthorized, "token has been revoked", "TOKEN_REVOKED")
                default:
                    abortWithError(c, http.StatusUnauthorized, "invalid token", "TOKEN_INVALID")
			}
			return
		}

		c.Set("employeeID", claims.EmployeeID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			abortWithError(c, http.StatusForbidden, "role not found in token", "ROLE_MISSING")
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			abortWithError(c, http.StatusForbidden, "invalid role format", "ROLE_INVALID")
			return
		}

		for _, r := range roles {
			if roleStr == r {
				c.Next()
				return
			}
		}

		abortWithError(c, http.StatusForbidden, "access denied", "ROLE_FORBIDDEN")
	}
}

func extractBearerToken(c *gin.Context) (string, bool) {
	header := c.GetHeader("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return "", false
	}
	token := strings.TrimPrefix(header, "Bearer ")
	if token == "" {
		return "", false
	}
	return token, true
}

func abortWithError(c *gin.Context, status int, message, code string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error": message,
		"code":  code,
	})
}