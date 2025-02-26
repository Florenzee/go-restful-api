package web

type CustomerUpdateRequest struct {
	CustomerID int    `validate:"required"`
	Name       string `validate:"required,max=200,min=1" json:"product_name"`
	Email      string `validate:"required,email" json:"email"`
	Phone      string `validate:"required,phone" json:"phone"`
	Address    string `validate:"required,address" json:"address"`
	LoyaltyPts int    `validate:"required,max=200" json:"loyalty_pts"`
}
