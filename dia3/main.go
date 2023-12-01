package main

import "main/ejercicio1"

func main() {
	println("----------- Ejercicio 1 -----------")

	newProduct := ejercicio1.Product{
		ID:          4,
		Name:        "Naranja",
		Price:       7,
		Description: "Naranja criolla",
		Category:    "Fruta",
	}

	println("------ Save ------")
	newProduct.Save()

	println("------ Get All ------")
	for _, product := range ejercicio1.Products {
		product.GetAll()
	}

	println("------ Get By ID ------")
	idToSearch := 2 //Es la manzana
	productSearch := ejercicio1.GetById(idToSearch)

	productSearch.GetAll()

	println("----------- Ejercicio 2 -----------")
}
