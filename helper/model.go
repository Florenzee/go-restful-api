package helper

import (
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

// category
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

// product
func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ProductID: product.ProductID,
		Name:      product.Name,
		Price:     product.Price,
		StockQty:  product.StockQty,
		Category:  product.Category,
	}
}

func ToProductResponses(product []domain.Product) []web.ProductResponse {
	var productResponse []web.ProductResponse
	for _, product := range product {
		productResponse = append(productResponse, ToProductResponse(product))
	}
	return productResponse
}

// customer
func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		LoyaltyPts: customer.LoyaltyPts,
	}
}

func ToCustomerResponses(customer []domain.Customer) []web.CustomerResponse {
	var customerResponse []web.CustomerResponse
	for _, customer := range customer {
		customerResponse = append(customerResponse, ToCustomerResponse(customer))
	}
	return customerResponse
}
