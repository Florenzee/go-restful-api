package web

type ProductCreateRequest struct {
	Name     string  `validate:"required,min=1,max=100" json:"name"`
	Price    float64 `validate:"required,min=1,max=100" json:"price"`
	StockQty int     `validate:"required,min=1,max=100" json:"stock_qty"`
	Category string  `validate:"required,min=1,max=100" json:"category"`
}
