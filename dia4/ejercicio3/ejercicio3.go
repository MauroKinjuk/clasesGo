/*
Ejercicio 3 - Impuestos de salario #3
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en
reemplazo de “Error()”, se implemente “errors.New()”
*/

package ejercicio3

import "errors"

type SalaryError struct{}

var ErrSalary = errors.New("Error: salary is less than 10000")

func CheckSalary(salary int) error {

	if salary < 10000 {
		return ErrSalary
	}

	return nil
}
