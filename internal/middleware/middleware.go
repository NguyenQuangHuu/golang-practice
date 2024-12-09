package middleware

import (
	"awesomeProject/internal/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequireAuthentication(c *gin.Context) {
	tokenString, err := helpers.GetToken(c)
	if err != nil {
		helpers.HandleUnauthorized(c, err)
		return
	}
	err = helpers.VerifyToken(tokenString)
	if err != nil {
		helpers.HandleUnauthorized(c, err)
		return
	}
	c.Next()
}

func RoleRequired(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := helpers.GetToken(c)
		if err != nil {
			helpers.HandleUnauthorized(c, err)
			return
		}
		claims, err := helpers.GetClaims(tokenString)
		if err != nil {
			helpers.HandleUnauthorized(c, err)
			return
		}
		roles := claims["roles"]
		if !helpers.HasRole(roles, role) {
			helpers.HandleForbidden(c, fmt.Errorf("no permission to access this resource"))
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}
