package web

type CustomerResponse struct {
	CustomerID int    `json:"cust_id"`
	Name       string `json:"cust_name"`
	Email      string `json:"cust_email"`
	Phone      string `json:"cust_phone"`
	Address    string `json:"cust_address"`
	LoyaltyPts int    `json:"loyalty_pts"`
}
