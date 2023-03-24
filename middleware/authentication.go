package middleware

import (
	"fmt"
	"strings"
	"testproject/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizationMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")
		if err := validateToken(token); err != nil {
			service.Resp(c, 401, "Invalid token", nil)
			return
		}
	}
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("MySignature"), nil
	})
	return err
}
