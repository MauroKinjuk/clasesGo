/*
Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el
mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza
el mínimo imponible (el mensaje mostrado por consola deberá decir: “Error: the minimum taxable amount is 150,000
and the salary entered is: [salary]”, siendo [salary]
el valor de tipo int pasado por parámetro).
*/

package ejercicio4

import "fmt"

type SalaryError struct{}

func CheckSalary(salary int) error {

	if salary < 10000 {
		return fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)
	}

	return nil
}
