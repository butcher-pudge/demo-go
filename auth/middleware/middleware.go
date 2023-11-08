package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-learning-demo/auth/constants"
	authErrors "go-learning-demo/auth/error"
	"go-learning-demo/auth/service"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	authService service.IAuthService
}

func NewAuthMiddleware(authService service.IAuthService) gin.HandlerFunc {
	return (&AuthMiddleware{
		authService: authService,
	}).Handle
}

func (middleware *AuthMiddleware) Handle(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := middleware.authService.ParseToken(context.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authErrors.ErrInvalidAccessToken) {
			status = http.StatusUnauthorized
		}

		context.AbortWithStatus(status)
		return
	}

	context.Set(constants.CtxUserKey, user)
}
