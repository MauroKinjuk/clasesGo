/*
Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo
“int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: salary is less than 10000" y lanzalo en caso de que “salary” sea menor o igual
a 10000. La validación debe ser hecha con la función “Is()” dentro del “main”.
*/

package ejercicio2

type SalaryError struct{}

func (e SalaryError) Error() string {
	return "Error: salary is less than 10000"
}

var ErrSalary = &SalaryError{}

func CheckSalary(salary int) error {
	if salary < 10000 {
		return ErrSalary
	}

	return nil
}
