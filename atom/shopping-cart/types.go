package shopping_cart

type Cart struct {
	ID int `json:"id"`
	UserID    int `json:"user_id"`
    ProductID int `json:"product_id"`
    Quantity  int `json:"quantity"`
}

type AddCartRequest struct {
	UserID    int `json:"user_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,gt=0"`
}

type GetCartRequest struct{
	UserID int `json:"user_id" validate:"required"`
}

type GetCartResponse struct {
	ID int `json:"id"`
	UserID    int `json:"user_id"`
	Username string `json:"username"`
	ProductID int `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity  int `json:"quantity"`
}

type DeleteCartRequest struct{
	UserID    int `json:"user_id" validate:"required"`
	CartID int `json:"cart_id" validate:"required"`
}