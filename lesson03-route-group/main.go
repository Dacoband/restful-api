package main

import (
	v1handler "dacoband.com/golang/internal/api/v1/handler"
	v2handler "dacoband.com/golang/internal/api/v2/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		userHandlerV1 := v1handler.NewUserHandler()
		user.GET("/users", userHandlerV1.GetUserV1)
		user.GET("/:id", userHandlerV1.GetUserByIdV1)
		user.POST("/", userHandlerV1.PostUserV1)
		user.PUT("/:id", userHandlerV1.PutUserV1)
		user.DELETE("/:id", userHandlerV1.DeleteUserV1)

		product := v1.Group("/product")
		productHandlerV1 := v1handler.NewProductHandler()
		product.GET("/", productHandlerV1.GetProductV1)
		product.GET("/:id", productHandlerV1.GetProductByIdV1)
		product.POST("/", productHandlerV1.PostProductV1)
		product.PUT("/:id", productHandlerV1.PutProductV1)
		product.DELETE("/:id", productHandlerV1.DeleteProductV1)
	}
	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/user")
		userHandlerV2 := v2handler.NewUserHandler()
		user.GET("/users", userHandlerV2.GetUserV2)
		user.GET("/:id", userHandlerV2.GetUserByIdV2)
		user.POST("/", userHandlerV2.PostUserV2)
		user.PUT("/:id", userHandlerV2.PutUserV2)
		user.DELETE("/:id", userHandlerV2.DeleteUserV2)
	}
	r.Run(":8080")
}
