/*
Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al
pegarle responda un texto que diga “pong”
1. El endpoint deberá ser de método GET
2. La respuesta de “pong” deberá ser enviada como texto, NO como JSON
*/

package ejercicio_1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ejercicio1() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
