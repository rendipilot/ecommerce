package orders

type Order struct {
	ID         int     `json:"id"`
    UserID     int     `json:"user_id"`
    TotalPrice float64 `json:"total_price"`
}

type OrderItem struct {
	ID        int     `json:"id"`
    OrderID   int     `json:"order_id"`
    ProductID int     `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type CartsItem struct {
	ProductID int     `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

type OrderRequest struct {
	UserID     int     `json:"user_id" validate:"required"`
}

type OrderResponse struct {
	OrderID int `json:"order_id"`
}