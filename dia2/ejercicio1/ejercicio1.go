/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y
si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

package ejercicio1

func CalculoImpuesto(salario float64) (impuesto int) {

	porcentaje := 17

	if salario > 150000 {
		impuesto = porcentaje + 10
	} else if salario > 50000 && salario <= 150000 {
		impuesto = porcentaje
	}

	return impuesto
}
