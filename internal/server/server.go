package server

import (
	"log"

	er "mp/internal/errors"
	h "mp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ServerInit() {
	r := gin.Default()
	r.POST("/user", h.NewUser)
	if err := r.Run(); err != nil {
		log.Println(er.ServerInitErr, err)
		return
	}

}
