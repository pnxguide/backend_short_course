package form

type Product struct {
	ProductID       int     `json:product_id`
	ProductName     string  `json:product_name`
	ProductQuantity int     `json:product_quantity`
	ProductPrice    float64 `json:product_price`
}
