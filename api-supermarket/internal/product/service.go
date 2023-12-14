package product

import "api-supermarket/internal/domain"

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{repo}
}

func (s ProductService) GetAll() []domain.Product {
	return s.repository.GetAll()
}

func (s ProductService) GetById(id int64) domain.Product {
	return s.repository.GetById(id)
}

func (s ProductService) Create(product domain.ProductCreate) domain.Product {
	return s.repository.Create(product)
}
