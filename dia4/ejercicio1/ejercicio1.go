/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo
“int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: the salary entered does not reach the taxable minimum" y lanzalo en caso de
que “salary” sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el
mensaje “Must pay tax”.
*/

package ejercicio1

type SalaryError struct {
	Message string
}

func (e SalaryError) Error() string {
	return e.Message
}

func CalcularImpuesto(salary int) error {
	if salary < 150000 {
		return SalaryError{"Error: the salary entered does not reach the taxable minimum"}
	}
	return nil
}
