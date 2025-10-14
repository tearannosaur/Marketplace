package handlers

import (
	"log"
	m "mp/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	var user m.UserRequest
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("Ошибка создания пользователя", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка создания пользователя",
		})
		return
	}
	newUser, err := m.CreateUser(user)
	if err != nil {
		log.Println("Ошибка создания пользователя", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка создания пользователя",
		})
		return
	}
	UserResponse := m.UserResponse{
		User_id: newUser.User_id,
		Login:   newUser.Login,
	}
	log.Println("Пользователь успешно создан\nЛогин:", newUser.Login, "\nId:", newUser.User_id)
	c.JSON(http.StatusCreated, gin.H{
		"success": UserResponse,
	})
}
