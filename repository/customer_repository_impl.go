package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db: db}
}

// save customer
func (repository *CustomerRepositoryImpl) Save(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	if err := repository.db.WithContext(ctx).Create(&customer).Error; err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

// update customer
func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	if err := repository.db.WithContext(ctx).Save(&customer).Error; err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

// delete product
func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, customer domain.Customer) error {
	if err := repository.db.WithContext(ctx).Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}

// FindById - get customer by ID
func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, customerId int) (domain.Customer, error) {
	var customer domain.Customer
	err := repository.db.WithContext(ctx).First(&customer, customerId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customer, errors.New("customer not found")
	}
	return customer, err
}

// FindAll - Get all customer
func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) ([]domain.Customer, error) {
	var customers []domain.Customer
	err := repository.db.WithContext(ctx).Find(&customers).Error
	return customers, err
}
