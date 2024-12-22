package helpers

import (
	"awesomeProject/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// SecretKey for generate token
var SecretKey = []byte("very-secret-key")

// GenerateToken is used to generate token code by username, algorithm signed method and secret key
func GenerateToken(user *model.User, expỉredTime time.Duration) (string, error) {
	claim := jwt.MapClaims{
		"iss":   "go-lang-server",
		"sub":   user.Username,
		"aud":   "angular-client",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(expỉredTime).Unix(),
		"roles": user.Roles,
	}
	tokenGenerated := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := tokenGenerated.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

// VerifyToken use to verify token given by client side
func VerifyToken(tokenString string) error {
	claims, err := GetClaims(tokenString)
	if err != nil {
		return err
	}
	if err = TokenExpired(claims); err != nil {
		return err
	}

	return nil
}

func GetClaims(token string) (jwt.MapClaims, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := claim.Claims.(jwt.MapClaims)
	if !ok || !claim.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	} else {
		return claims, nil
	}
}

func TokenExpired(claim jwt.MapClaims) error {
	expiration, ok := claim["exp"].(float64)
	if !ok {
		return fmt.Errorf("invalid token: missing experation")
	}
	if time.Unix(int64(expiration), 0).Before(time.Now()) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func GetToken(c *gin.Context) (string, error) {
	tokenString, err := c.Cookie("token-jwt")
	if err != nil {
		tokenString = c.Request.Header.Get("Authorization")
		if strings.Contains(tokenString, "Bearer ") && len(tokenString) > 7 {
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		}
	}
	if tokenString == "" {
		return "", fmt.Errorf("token is required")
	}

	return tokenString, nil
}

func HasRole(role interface{}, requiredRole []string) bool {
	roleMap := make(map[string]struct{})
	if roleSlice, ok := role.([]interface{}); ok {
		for _, role := range roleSlice {
			roleMap[role.(string)] = struct{}{}
		}
	}
	for _, role := range requiredRole {
		if _, ok := roleMap[role]; ok {
			return true
		}
	}
	return false
}

func GetUsername(c *gin.Context) (string, error) {
	tokenString, err := GetToken(c)
	if err != nil {
		return "", err
	}
	claims, err := GetClaims(tokenString)
	if err != nil {
		return "", err
	}
	username := claims["sub"].(string)
	return username, nil
}
