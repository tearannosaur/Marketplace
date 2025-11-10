package server

import (
	"log"

	er "mp/internal/errors"
	h "mp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ServerInit(h *h.HandlerModule) {
	r := gin.Default()
	r.POST("/user/signup", h.SignUp)
	r.POST("/user/login", h.Login)
	if err := r.Run(); err != nil {
		log.Println(er.ServerInitErr, err)
		return
	}
}
