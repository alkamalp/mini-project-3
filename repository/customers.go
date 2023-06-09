package repository

import (
	"github.com/alkamalp/crm-golang/entity"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}
}

type CustomerInterfaceRepo interface {
	CreateCustomer(Customer *entity.Customer) (*entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	UpdateCustomer(customer *entity.Customer, id uint) (any, error)
	DeleteCustomer(id uint) (any, error)
}

// CreateCustomer new Customer
func (repo Customer) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Create(customer).Error
	return customer, err
}

// GetCustomerById get single Customer by id
func (repo Customer) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	repo.db.First(&customer, "id = ? ", id)
	return customer, nil
}

// UpdateCustomer multiple fields
func (repo Customer) UpdateCustomer(customer *entity.Customer, id uint) (any, error) {
	err := repo.db.Model(&entity.Customer{}).Where("id = ?", id).
		Updates(customer).Error
	return nil, err
}

// DeleteCustomer by Id and email
func (repo Customer) DeleteCustomer(id uint) (any, error) {
	err := repo.db.Model(&entity.Customer{}).
		Where("id = ?", id).
		Delete(&entity.Customer{}).
		Error
	return nil, err
}
