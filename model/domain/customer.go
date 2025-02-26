package domain

type Customer struct {
	CustomerID int    `gorm:"primaryKey" column:"id"`
	Name       string `gorm:"column:customer_name"`
	Email      string `gorm:"column:customer_email"`
	Phone      string `gorm:"column:customer_phone"`
	Address    string `gorm:"column:customer_address"`
	LoyaltyPts int    `gorm:"column:loyalty_pts"`
}
