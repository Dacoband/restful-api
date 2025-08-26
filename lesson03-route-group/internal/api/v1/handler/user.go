package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List users v1",
	})
}
func (u *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID v1",
	})
}
func (u *UserHandler) PostUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create user v1",
	})
}
func (u *UserHandler) PutUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update user v1",
	})
}
func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete user v1",
	})
}
