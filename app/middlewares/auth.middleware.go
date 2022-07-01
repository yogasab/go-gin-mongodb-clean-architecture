package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-gin-mongodb-clean-architecture/app/services/auth"
	"go-gin-mongodb-clean-architecture/app/services/user"
	"net/http"
	"strings"
)

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := ctx.GetHeader("Authorization")

		if !strings.Contains(headers, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Failed to authenticate user", "status": "error", "errors": "Invalid token format"})
			return
		}

		splittedHeaders := strings.Split(headers, " ")
		if splittedHeaders[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Failed to authenticate user", "status": "error", "errors": "Invalid token format"})
			return
		}

		jwtToken := splittedHeaders[1]
		token, err := authService.ValidateToken(jwtToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Failed to authenticate user", "status": "error", "errors": "Invalid token format"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Failed to authenticate user", "status": "error", "errors": "Invalid token format"})
			return
		}

		user_id := claims["user_id"]
		registeredUser, err := userService.GetUserByID(fmt.Sprintf("%s", user_id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Failed get registered user", "status": "error", "errors": err.Error()})
			return
		}

		ctx.Set("user", registeredUser)
	}
}
