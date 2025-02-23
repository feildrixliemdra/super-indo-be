package payload

type CreateCartItemRequest struct {
	ProductID uint64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type GetAllCartItemResponse struct {
	CartID uint64 `json:"cart_id"`
	Items  []Item `json:"items"`
}

type Item struct {
	ID           uint64 `json:"id"`
	ProductID    uint64 `json:"product_id"`
	Quantity     int    `json:"quantity"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
	ProductImage string `json:"product_image"`
	CategoryName string `json:"category_name"`
	CategoryID   uint64 `json:"category_id"`
}
