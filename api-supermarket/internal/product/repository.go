package product

import (
	"api-supermarket/internal/domain"
	"errors"
)

type ProductRepository struct {
	productsDB []domain.Product
	nextID     int
}

func NewProductRepository(prods []domain.Product) ProductRepository {
	return ProductRepository{
		productsDB: prods,
		nextID:     len(prods) + 1}
}

func (r ProductRepository) GetAll() []domain.Product {
	return r.productsDB
}

func (r ProductRepository) GetById(id int64) domain.Product {
	return r.productsDB[id-1]
}

func checkExpiration(expiration string) bool {
	// Check if Date is in format XX/XX/XXXX
	if len(expiration) != 10 {
		return false
	}
	// Check if Date has the correct format
	if expiration[2] != '/' || expiration[5] != '/' {
		return false
	}
	// Check if Day is valid
	if expiration[0] < '0' || expiration[0] > '3' || expiration[1] < '0' || expiration[1] > '9' {
		return false
	}
	// check if Month is valid
	if expiration[3] < '0' || expiration[3] > '1' || expiration[4] < '0' || expiration[4] > '9' {
		return false
	}
	// Check if Year is valid
	if expiration[6] < '0' || expiration[6] > '9' || expiration[7] < '0' || expiration[7] > '9' || expiration[8] < '0' || expiration[8] > '9' || expiration[9] < '0' || expiration[9] > '9' {
		return false
	}
	return true
}

func checkCodeUnique(code string, products []domain.Product) bool {
	for _, product := range products {
		if product.CodeValue == code {
			return false
		}
	}
	return true
}

func (r *ProductRepository) Create(product domain.ProductCreate) (domain.Product, error) {
	/*
		DONE - No es necesario pasar el Id, al momento de añadirlo se debe inferir del estado de la lista de productos, verificando que no se repitan ya que debe ser un campo único.
		DONE - Ningún dato puede estar vacío, exceptuando is_published (vacío indica un valor false).
		DONE - El campo code_value debe ser único para cada producto.
		Los tipos de datos deben coincidir con los definidos en el planteo del problema.
		DONE - La fecha de vencimiento debe tener el formato: XX/XX/XXXX, además debemos verificar que día, mes y año sean valores válidos.
		DONDE - Recordá: si una consulta está mal formulada por parte del cliente, el status code cae en los 4XX.
	*/

	// If data is empty, return error
	if product.Name == "" || product.Quantity == 0 || product.CodeValue == "" || product.Expiration == "" || product.Price == 0 {
		return domain.Product{}, errors.New("data is empty")
	}

	// Check if code_value is unique
	if !checkCodeUnique(product.CodeValue, r.productsDB) {
		return domain.Product{}, errors.New("code_value is not unique")
	}

	// Check if date is valid in format XX/XX/XXXX
	if !checkExpiration(product.Expiration) {
		return domain.Product{}, errors.New("expiration date is invalid")
	}

	newProduct := domain.Product{
		Id:          r.nextID,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}

	r.productsDB = append(r.productsDB, newProduct)
	r.nextID++

	return newProduct, nil
}
