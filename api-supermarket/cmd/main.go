package main

import (
	"api-supermarket/cmd/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	server := gin.Default()

	// Group the routes
	productsGroup := server.Group("/products")
	productRouter := handler.NewProductRouter(productsGroup)
	productRouter.Routes()

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
