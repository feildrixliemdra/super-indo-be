package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenExpired = 7 * 24 * time.Hour //7 days expired
)

type JWTUser struct {
	UserID uint64
	Email  string
}

// ValidateJWT validate jwt token from header
func ValidateJWT(c *gin.Context, secretKey string) error {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// ExtractToken get jwt token string from Authorization header
func ExtractToken(c *gin.Context) string {
	// expected header key:
	// Authorization: Bearer {token}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func GenerateJWT(user JWTUser, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     time.Now().Add(TokenExpired).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func ParseJWT(tokenString string, secretKey string) (JWTUser, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return JWTUser{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return JWTUser{}, fmt.Errorf("invalid token claims")
	}

	userID := uint64(claims["user_id"].(float64))
	email := claims["email"].(string)

	return JWTUser{UserID: userID, Email: email}, nil
}
