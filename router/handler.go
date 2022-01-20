package router

import (
	"github.com/gin-gonic/gin"
	"gomasters/service"
)

type Handler struct {
}

func (h *Handler)InitRoutes() *gin.Engine{
	router := gin.New()

	users := router.Group("/users")
	{
		users.PUT("/:id", h.editUser)
		users.POST("/", h.createUser)
		users.GET("/:id", h.getUser)
	}
	return router
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{}
}
