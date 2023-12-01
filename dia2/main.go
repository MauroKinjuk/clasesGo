package main

import (
	"example/hello/Desktop/dia2/ejercicio1"
	"example/hello/Desktop/dia2/ejercicio2"
	"example/hello/Desktop/dia2/ejercicio3"
	"fmt"
)

func main() {
	fmt.Println("Hola!")
	// Ejercicio 1
	fmt.Println("El impuesto seria de: ", ejercicio1.CalculoImpuesto(160000))

	//Ejercicio 2
	fmt.Println("El promedio de notas son: ", ejercicio2.PromedioNotas(1, 2, 3))

	//Ejercicio 3
	fmt.Println("El salario es: ", ejercicio3.CalcSalario(10, "B"))

}
