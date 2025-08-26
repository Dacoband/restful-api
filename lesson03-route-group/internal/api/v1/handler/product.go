package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List products v1",
	})
}
func (p *ProductHandler) GetProductByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID v1",
	})
}
func (p *ProductHandler) PostProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create product v1",
	})
}
func (p *ProductHandler) PutProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update product v1",
	})
}
func (p *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete product v1",
	})
}
