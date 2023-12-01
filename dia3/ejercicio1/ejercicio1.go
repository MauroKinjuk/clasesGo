/*
										Ejercicio 1
Crear un programa que cumpla los siguiente puntos:
1. Tener una estructura llamada Product con los campos ID, Name, Price,
Description y Category.
2. Tener un slice global de Product llamado Products instanciado con valores.
3. 2 métodos asociados a la estructura Product: Save(), GetAll(). El método
Save() deberá tomar el slice de Products y añadir el producto desde el cual
se llama al método. El método GetAll() deberá imprimir todos los productos
guardados en el slice Products.
4. Una función getById() al cual se le deberá pasar un INT como parámetro y
retorna el producto correspondiente al parámetro pasado.
5. Ejecutar al menos una vez cada método y función definido desde main().
*/

package ejercicio1

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{ID: 1, Name: "Banana", Price: 10.0, Description: "Banana amarilla", Category: "Fruta"},
	{ID: 2, Name: "Manzana", Price: 8.0, Description: "Manzana verde", Category: "Fruta"},
	{ID: 3, Name: "Ajo", Price: 2.0, Description: "Ajo entero", Category: "Verdura"},
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p Product) GetAll() {
	fmt.Printf("ID: %d, Name: %s, Price: %2.f, Description: %s, Category: %s \n", p.ID, p.Name, p.Price, p.Description, p.Category)
}

func getById(id int) Product {

	return Product{}
}
