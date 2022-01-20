package router

import (
	"github.com/gin-gonic/gin"
	"gomasters/models"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context){
	var user models.User
	if err := ctx.BindJSON(&user); err != nil{
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreateUser(user)
	if err != nil{
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated,  map[string]int{
		"id":id,
	})
}
func (h *Handler) editUser(ctx *gin.Context){

}
func (h *Handler) getUser(ctx *gin.Context){

}