package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Setting headers
	gin.SetMode(gin.ReleaseMode) // Set Gin to release mode for production

	r := gin.Default()
	// Set up a route
	// This route will respond to GET requests on /demo
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"users": []string{"Alice", "Bob", "Charlie"}}) // endpoint for user list
	})

	r.GET("/users/:user_id", func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		ctx.JSON(200, gin.H{"user": userID})
	})

	r.GET("/user/:user_id", func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")
		ctx.JSON(200, gin.H{
			"Thông tin của user": user_id,
			"Thông tin bổ sung":  "Some additional info about user " + user_id,
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"products": []string{"Laptop", "Smartphone", "Tablet"}})
	})

	r.GET("/product/detail/:product_name", func(ctx *gin.Context) {
		product_name := ctx.Param("product_name")
		price := ctx.Query("price")
		color := ctx.Query("color")

		ctx.JSON(200, gin.H{
			"Thông tin của sản phẩm": product_name,
			"Thông tin bổ sung":      "Some additional info about product " + product_name,
			"Giá":                    price,
			"Màu sắc":                color,
		})
	})
	r.Run(":8080")
}
