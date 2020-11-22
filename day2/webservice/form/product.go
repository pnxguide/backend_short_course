package form

type Product struct {
	ProductID       uint64  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductQuantity uint64  `json:"product_quantity"`
	ProductPrice    float64 `json:"product_price"`
}
