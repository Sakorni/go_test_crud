package router

import (
	"github.com/gin-gonic/gin"
	"gomasters/models"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(ctx *gin.Context){
	user,err := scanUser(ctx)
	if err != nil{
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
	user,err := scanUser(ctx)
	if err != nil{
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	strId, ok := ctx.Params.Get("id")
	if !ok{
		newErrorResponse(ctx, http.StatusBadRequest, "id must be provided in request params")
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil || id < 0{
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id in params")
		return
	}

	user.ID = uint(id)
	err = h.service.UpdateUser(user)
	if err != nil{
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
}
func (h *Handler) getUser(ctx *gin.Context){
	paramId, ok := ctx.Params.Get("id")
	if !ok{
		newErrorResponse(ctx, http.StatusBadRequest, "no id provided")
		return
	}
	id, err := strconv.Atoi(paramId)
	if err != nil{
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id")
		return
	}
	user, err := h.service.GetUser(id)
	if err != nil{
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func scanUser(ctx *gin.Context) (models.User, error){
	var user models.User
	err := ctx.BindJSON(&user)
	return user, err
}