package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHell(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hola, como estas?")
}
func sayHelloName(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hola %s, como estas?", ctx.Param("name"))
}

func sayHelloConf(ctx *gin.Context) {
	query := ctx.Query("confidence")

	if query == "si" {
		ctx.String(http.StatusOK, "Hola Crack!")
	} else {
		ctx.String(http.StatusOK, "Hola %s.", ctx.Param("name"))
	}

}

func main() {
	//ejercicio_1.Ejercicio1()
	//ejercicio_2.Ejercicio2()

	server := gin.Default()

	sayHello := server.Group("/sayHello")

	sayHello.GET("/hello", sayHell)            //http://localhost:8080/sayHello/hello
	sayHello.GET("/:name", sayHelloName)       //http://localhost:8080/sayHello/Mauro
	sayHello.GET("/query/:name", sayHelloConf) //http://localhost:8080/sayHello/query/Mauro?confidence=si

	server.Run(":8080")
}
