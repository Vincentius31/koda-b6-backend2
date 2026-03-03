package repository

import "koda-b6-backend2/internal/models"

var DataProduct []models.Product

type ProductRepository struct{
	db *[]models.Product
}

func NewProductRepository(db *[]models.Product) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Get all Product
func (r *ProductRepository) GetAll() *[]models.Product{
	return r.db
}

// Get Product by Id
func (r *ProductRepository) GetById(id int) (*models.Product, int){
	for i, p := range *r.db {
		if p.Id == id {
			return &(*r.db)[i],i
		}
	}
	return nil, -1
}

// Create Product
func (r *ProductRepository) Create(product models.Product){
	*r.db = append(*r.db, product)
}

// Update Product
func (r *ProductRepository) Update(index int, updatedProduct models.Product){
	(*r.db)[index] = updatedProduct
}

// Delete Product
func (r *ProductRepository) Delete(index int){
	*r.db = append((*r.db)[:index], (*r.db)[:index+1]...)
}