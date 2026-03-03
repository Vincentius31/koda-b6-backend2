package service

import (
	"koda-b6-backend2/internal/models"
	"koda-b6-backend2/internal/repository"
)

type ProductService struct{
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository)*ProductService{
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAll() []models.Product {
	return *s.repo.GetAll()
}

func (s *ProductService) GetByID(id int) *models.Product {
	product, _ := s.repo.GetById(id)
	return product
}

func (s *ProductService) Create(req models.CreateProductRequest) {
	products := *s.repo.GetAll()
	newID := 1
	if len(products) > 0 {
		newID = products[len(products)-1].Id + 1
	}

	newProduct := models.Product{
		Id:    newID,
		Name:  req.Name,
		Price: req.Price,
	}
	s.repo.Create(newProduct)
}

func (s *ProductService) Update(id int, req models.UpdateProductRequest) bool {
	product, index := s.repo.GetById(id)
	if product == nil {
		return false
	}
	product.Name = req.Name
	product.Price = req.Price
	s.repo.Update(index, *product)
	return true
}

func (s *ProductService) Delete(id int) bool {
	_, index := s.repo.GetById(id)
	if index == -1 {
		return false
	}
	s.repo.Delete(index)
	return true
}