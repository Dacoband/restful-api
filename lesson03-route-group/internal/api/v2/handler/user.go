package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List users v2",
	})
}
func (u *UserHandler) GetUserByIdV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID v2",
	})
}
func (u *UserHandler) PostUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create user v2",
	})
}
func (u *UserHandler) PutUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user v2",
	})
}
func (u *UserHandler) DeleteUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete user v2",
	})
}
