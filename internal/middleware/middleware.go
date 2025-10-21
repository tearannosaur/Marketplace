package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth(c *gin.Context) {

	token, err := c.Cookie("Authorization")
	if err != nil {
		log.Println("Authorization error", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "missing or invalid token cookie",
		})
		return
	}
	claims, err := ValidateToken(token)
	if err != nil {
		log.Println("Token Validate error", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid or expired token",
		})
		return
	}
	userId, ok := claims["user_id"].(string)
	if !ok {
		log.Println("Token Validate error", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token payload",
		})
		return
	}
	user, err := uuid.Parse(userId)
	if err != nil {
		log.Println("UUID token parse error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}
	c.Set(`user`, user)
	c.Next()
}
