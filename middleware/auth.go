package middleware

import (
	"fmt"
	"strings"

	"github.com/alkamalp/crm-golang/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) {

	// Token yang diterima
	receivedToken := c.GetHeader("Authorization")
	signedToken := strings.Split(receivedToken, " ")

	// Verifikasi token dengan kunci rahasia
	token, err := jwt.Parse(signedToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		c.JSON(401, dto.DefaultErrorInvalidDataWithMessage("token tidak valid"))
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Token valid, akses klaim-klaim yang ada
		c.Set("Role", claims["sub"])
	} else {
		c.JSON(401, dto.DefaultErrorInvalidDataWithMessage("token tidak valid"))
		c.Abort()
		return
	}
	c.Next()
}
