package web

type CustomerCreateRequest struct {
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Email      string `validate:"required,email" json:"cust_email"`
	Phone      string `validate:"required,phone" json:"cust_phone"`
	Address    string `validate:"required,address" json:"cust_address"`
	LoyaltyPts int    `validate:"required,min=1,max=100" json:"loyalty_pts"`
}
