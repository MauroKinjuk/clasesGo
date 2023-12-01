/*Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la
cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes,
la categoría y que devuelva su salario.
*/

package ejercicio3

import (
	"strings"
)

const categoryC float64 = 1000
const categoryB float64 = 1500
const categoryA float64 = 3000

func CalcSalario(minutes int, category string) (salary float64) {

	category = strings.ToLower(category)
	hours := float64(minutes) / 60

	switch category {
	case "a":
		salary = hours * categoryA
		salary += salary * 0.5
	case "b":
		salary = hours * categoryB
		salary += salary * 0.2
	case "c":
		salary = hours * categoryC

	}

	return salary
}
