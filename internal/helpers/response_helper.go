package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleUnauthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}

func HandleForbidden(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
}
