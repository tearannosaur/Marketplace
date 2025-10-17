package handlers

import (
	"log"
	er "mp/internal/errors"
	m "mp/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerModule) NewUser(c *gin.Context) {
	var user m.UserRequest
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(er.IncorrectUserDataErr, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er.IncorrectUserDataErr,
		})
		return
	}
	newUser, err := m.CreateUser(user)
	if err != nil {
		log.Println(er.UserCreateErr, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.UserCreateErr,
		})
		return
	}
	//сохранить в бд
	UserResponse := m.UserResponse{
		UserId: newUser.UserId,
		Login:  newUser.Login,
	}
	log.Println("User successfully created\nLogin:", newUser.Login, "\nId:", newUser.UserId)
	c.JSON(http.StatusCreated, gin.H{
		"success": UserResponse,
	})

}
