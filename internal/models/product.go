package models

type Product struct{
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductRequest struct{
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductRequest struct{
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
