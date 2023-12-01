/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.
*/

package ejercicio2

import (
	"fmt"
)

func PromedioNotas(notas ...int) (promedio float64) {

	var (
		suma   int
		tamano int
	)

	tamano = len(notas)

	for _, nota := range notas {
		if nota < 0 {
			fmt.Println("error: no se permiten numeros negativos")
			return 0
		}
		suma += nota
	}

	return float64(suma) / float64(tamano)
}
