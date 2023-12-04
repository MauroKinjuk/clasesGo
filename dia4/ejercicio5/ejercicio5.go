/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
1. Desarrollá las funciones necesarias para permitir a la empresa calcular:
a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
- La función recibirá las horas trabajadas en el mes y el valor de la hora como
argumento.
- Dicha función deberá retornar más de un valor (salario calculado y error).
- En caso de que el salario mensual sea igual o superior a $150.000, se le
deberá descontar el 10 % en concepto de impuesto.
- En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o
un número negativo, la función debe devolver un error. El mismo tendrá que
indicar “Error: the worker cannot have worked less than 80 hours per month”.
*/

package ejercicio5

import "fmt"

func Salary(hours float64, costHour float64) (salary float64, err error) {
	const maxSalary = 150_000
	salary = hours * costHour

	if hours < 0 || hours < 80 {
		err = fmt.Errorf("the worker cannot have worked less than 80 hours per month")
	}

	if salary >= maxSalary {
		salary = salary - (salary * 0.1)
	}

	return salary, err
}
