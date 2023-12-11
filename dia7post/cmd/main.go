package main

import (
	"app/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	//app
	// - dependencies
	hd := handler.NewDefaultMovies()
	// - router
	rt := gin.Default()
	// - endpoints
	grMovie := rt.Group("/movies")
	grMovie.POST("/", hd.CreateMovie())

	rt.Run(":8080")
}
