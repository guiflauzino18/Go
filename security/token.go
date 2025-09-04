package security

import (
	"errors"
	"fmt"
	"go-project/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Generate Token
func TokenGenerate(userID uint64, username string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 8).Unix()
	permissions["userId"] = userID
	permissions["username"] = username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func TokenValidate(c *gin.Context) (jwt.MapClaims, error) {
	r := c.Request
	tokenString := extractTokenFromRequestHeader(r)
	token, err := jwt.Parse(tokenString, getSecretKey)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return claims, nil
	}

	return nil, errors.New("Token inválido")
}

func extractTokenFromRequestHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Métodos de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
