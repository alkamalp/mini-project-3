package customers

import (
	"time"

	"github.com/alkamalp/crm-golang/entity"
	"github.com/alkamalp/crm-golang/repository"
)

type UseCaseCustomer interface {
	CreateCustomer(customer CustomerParam) (entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	UpdateCustomer(customer CustomerParam, id uint) (any, error)
	DeleteCustomer(id uint) (any, error)
}

type useCaseCustomer struct {
	customerRepo repository.CustomerInterfaceRepo
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entity.Customer, error) {
	var newCustomer *entity.Customer

	newCustomer = &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc useCaseCustomer) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	customer, err := uc.customerRepo.GetCustomerById(id)
	return customer, err
}

func (uc useCaseCustomer) UpdateCustomer(customer CustomerParam, id uint) (any, error) {
	var editCustomer *entity.Customer
	editCustomer = &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		UpdatedAt:  time.Now(),
	}

	_, err := uc.customerRepo.UpdateCustomer(editCustomer, id)
	if err != nil {
		return *editCustomer, err
	}
	return *editCustomer, nil
}

func (uc useCaseCustomer) DeleteCustomer(id uint) (any, error) {
	_, err := uc.customerRepo.DeleteCustomer(id)
	return nil, err
}
