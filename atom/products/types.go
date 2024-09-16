package products

type Products struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Category string `json:"category"`
	Stock int `json:"stock"`
}

type ProductRequest struct {
	Category string `json:"category"`
}