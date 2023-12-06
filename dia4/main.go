package main

import (
	"dia4/ejercicio1"
	"dia4/ejercicio2"
	"dia4/ejercicio3"
	"dia4/ejercicio4"
	"dia4/ejercicio5"
	"errors"
	"fmt"
)

func main() {
	/*-------------- Ejercicio 1 --------------*/
	fmt.Println("Ejercicio 1")
	salary := 160000

	err := ejercicio1.CalcularImpuesto(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Must pay tax")
	}

	/*-------------- Ejercicio 2 --------------*/
	fmt.Println("Ejercicio 2")

	salary2 := 9000

	err2 := ejercicio2.CheckSalary(salary2)

	//is
	if errors.Is(err2, ejercicio2.ErrSalary) {
		fmt.Println(err2)
	}

	/*-------------- Ejercicio 3 --------------*/
	fmt.Println("Ejercicio 3")

	salary3 := 9000

	err3 := ejercicio3.CheckSalary(salary3)

	if err3 != nil {
		fmt.Println(err3)
	}

	/*-------------- Ejercicio 4 --------------*/
	fmt.Println("Ejercicio 4")

	salary4 := 9000

	err4 := ejercicio4.CheckSalary(salary4)

	if err4 != nil {
		fmt.Println(err4)
	}

	/*-------------- Ejercicio 5 --------------*/
	fmt.Println("Ejercicio 5")

	hours := 79.0
	costHour := 10.0

	salary5, err5 := ejercicio5.Salary(hours, costHour)

	if err5 != nil {
		fmt.Println(err5)
		return
	} else {
		fmt.Println(salary5)
	}

}
