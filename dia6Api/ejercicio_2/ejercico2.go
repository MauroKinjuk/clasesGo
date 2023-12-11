/*
Ejercicio 2 - Manipulando el body
Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con
nombre y apellido que al pegarle deberá responder en texto “Hola + nombre + apellido”
1. El endpoint deberá ser de método POST
2. Se deberá usar el package JSON para resolver el ejercicio
3. La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
4. La estructura deberá ser como esta:
{
“nombre”: “Andrea”,
“apellido”: “Rivas” }
*/

package ejercicio_2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func Ejercicio2() {
	r := gin.Default()

	r.POST("/saludo", func(c *gin.Context) {
		// Decodifica el cuerpo de la solicitud JSON en la estructura SaludoRequest
		var saludoRequest User
		if err := c.ShouldBindJSON(&saludoRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Construye el mensaje de saludo
		saludo := "Hola " + saludoRequest.Name + " " + saludoRequest.LastName

		// Responde con el saludo como texto
		c.String(http.StatusOK, saludo)
	})

	r.Run(":8080")

}
