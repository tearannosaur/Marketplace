package handlers

import (
	"log"
	er "mp/internal/errors"
	middleware "mp/internal/middleware"
	m "mp/internal/models"
	u "mp/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerModule) SignUp(c *gin.Context) {
	var user m.UserRequest
	err := c.BindJSON(&user)
	if err != nil {
		log.Println(er.IncorrectJsonBody, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er.IncorrectJsonBody,
		})
		return
	}
	err = m.ValidateUserData(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er.IncorrectJsonBody,
		})
		return
	}
	newUser, err := m.CreateUser(user)
	if err != nil {
		log.Println(er.UserCreateErr, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.InternalServerErr,
		})
		return
	}
	err = h.repo.UserExist(newUser.Login)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{
			"error": er.UserAlreadyExistErr,
		})
		return
	}
	err = h.repo.SaveUser(newUser)
	if err != nil {
		log.Println(er.SaveUserDBErr, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.InternalServerErr,
		})
		return
	}
	UserResponse := m.UserResponse{
		UserId: newUser.UserId,
		Login:  newUser.Login,
	}
	log.Println("User successfully created\nLogin:", newUser.Login, "\nId:", newUser.UserId)
	c.JSON(http.StatusCreated, gin.H{
		"success": UserResponse,
	})

}

func (h *HandlerModule) Login(c *gin.Context) {
	//получаем данные
	var body m.UserRequest
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er.InvalidJsonBodyErr,
		})
		return
	}
	//проверяем на наличие логина
	user, err := h.repo.GetUserLogin(body.Login)
	if err != nil {
		log.Println("Get user error:", err)
		if err.Error() == er.UserDoesntExist {
			c.JSON(http.StatusNotFound, gin.H{"error": er.UserDoesntExist})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": er.InternalServerErr})
		}
		return
	}
	//проверяем пароль
	err = u.VerifyPassword(user.Password, body.Password)
	if err != nil {
		log.Println("Password error", user.Password, body.Password)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": er.InvalidPasswordErr,
		})
		return
	}
	//генерируем токен
	tokenString, err := middleware.GenerateToken(user)
	if err != nil {
		log.Println("Generate token error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": er.InternalServerErr,
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 12*60*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"user_id": user.UserId,
		"role":    user.Role,
		"token":   tokenString,
	})
	log.Println("User login:", user.UserId, user.Role)
}
