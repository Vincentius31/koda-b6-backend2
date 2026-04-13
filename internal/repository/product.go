package repository

import (
	"context"
	"koda-b6-backend2/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetById(ctx context.Context, id int) (*models.Product, error) {
	var product models.Product
	err := r.db.QueryRow(ctx, "SELECT id, name, price FROM products WHERE id = $1", id).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(ctx context.Context, product models.Product) error {
	_, err := r.db.Exec(ctx, "INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	return err
}

func (r *ProductRepository) Update(ctx context.Context, id int, updatedProduct models.Product) error {
	_, err := r.db.Exec(ctx, "UPDATE products SET name = $1, price = $2 WHERE id = $3", updatedProduct.Name, updatedProduct.Price, id)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	return err
}
