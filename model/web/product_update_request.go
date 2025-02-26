package web

type ProductUpdateRequest struct {
	ProductID int     `validate:"required"`
	Name      string  `validate:"required,max=200,min=1" json:"product_name"`
	Price     float64 `validate:"required,min=1" json:"product_price"`
	StockQty  int     `validate:"required,min=0" json:"stock_qty"`
	Category  string  `validate:"required,max=200" json:"category"`
}
