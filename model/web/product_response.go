package web

type ProductResponse struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"product_name"`
	Price     float64 `json:"product_price"`
	StockQty  int     `json:"stock_qty"`
	Category  string  `json:"category"`
}
