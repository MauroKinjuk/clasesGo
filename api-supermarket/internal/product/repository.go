package product

import "api-supermarket/internal/domain"

type ProductRepository struct {
	productsDB []domain.Product
}

func NewProductRepository(prods []domain.Product) ProductRepository {
	return ProductRepository{prods}
}

func (r ProductRepository) GetAll() []domain.Product {
	return r.productsDB
}

func (r ProductRepository) GetById(id int64) domain.Product {
	return r.productsDB[id-1]
}

func (r ProductRepository) Create(product domain.ProductCreate) domain.Product {
	/*
		No es necesario pasar el Id, al momento de añadirlo se debe inferir del estado de la
		lista de productos, verificando que no se repitan ya que debe ser un campo único.
		Ningún dato puede estar vacío, exceptuando is_published (vacío indica un valor false).
		El campo code_value debe ser único para cada producto.
		Los tipos de datos deben coincidir con los definidos en el planteo del problema.
		La fecha de vencimiento debe tener el formato: XX/XX/XXXX,
		además debemos verificar que día, mes y año sean valores válidos.
		Recordá: si una consulta está mal formulada por parte del cliente, el status code cae en los 4XX.
	*/
	newProduct := domain.Product{
		Id:          len(r.productsDB) + 1,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}

	r.productsDB = append(r.productsDB, newProduct)

	return newProduct
}
