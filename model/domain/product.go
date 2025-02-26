package domain

type Product struct {
	ProductID int     `gorm:"primaryKey" column:"id"`
	Name      string  `gorm:"column:product_name"`
	Price     float64 `gorm:"column:price"`
	StockQty  int     `gorm:"column:stock_qty"`
	Category  string  `gorm:"column:product_category"`
}
