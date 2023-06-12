package customers

import (
	"errors"
	"fmt"
	"github.com/alkamalp/crm-golang/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockCustomerRepo struct {
	mock.Mock
}

func (m *MockCustomerRepo) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	args := m.Called(customer)
	result := args.Get(0)
	err := args.Error(1)
	if result == nil {
		return nil, err
	}
	return result.(*entity.Customer), err
}

func TestCreateCustomer(t *testing.T) {

	mockRepo := new(MockCustomerRepo)

	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	mockRepo.On("CreateCustomer", mock.AnythingOfType("*entity.Customer")).Return(&entity.Customer{}, nil)

	createdCustomer, err := useCase.CreateCustomer(customer)

	mockRepo.AssertCalled(t, "CreateCustomer", mock.AnythingOfType("*entity.Customer"))

	assert.NotNil(t, createdCustomer)
	assert.NoError(t, err)
}

func (m *MockCustomerRepo) GetCustomerById(id uint) (entity.Customer, error) {
	args := m.Called(id)
	result := args.Get(0)
	err := args.Error(1)
	if result == nil {
		return entity.Customer{}, err
	}
	return result.(entity.Customer), err
}

func TestGetCustomerById(t *testing.T) {

	mockRepo := new(MockCustomerRepo)

	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	customerID := uint(1)
	customer := entity.Customer{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	mockRepo.On("GetCustomerById", customerID).Return(customer, nil)

	result, err := useCase.GetCustomerById(customerID)

	mockRepo.AssertCalled(t, "GetCustomerById", customerID)

	assert.Equal(t, customer, result)
	assert.NoError(t, err)
}

func TestGetCustomerById_Error(t *testing.T) {

	mockRepo := new(MockCustomerRepo)

	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	customerID := uint(1)
	expectedError := errors.New("failed to get customer")

	mockRepo.On("GetCustomerById", customerID).Return(entity.Customer{}, expectedError)

	result, err := useCase.GetCustomerById(customerID)

	mockRepo.AssertCalled(t, "GetCustomerById", customerID)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Customer{}, result)
}

func (m *MockCustomerRepo) UpdateCustomer(customer *entity.Customer, id uint) (interface{}, error) {
	args := m.Called(customer, id)
	return args.Get(0), args.Error(1)
}

func TestUpdateCustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepo)
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	customerID := uint(1)
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	expectedCustomer := &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		UpdatedAt:  time.Now(),
	}

	mockRepo.On("UpdateCustomer", expectedCustomer, customerID).Return(expectedCustomer, nil)
	result, err := useCase.UpdateCustomer(customer, customerID)
	mockRepo.AssertCalled(t, "UpdateCustomer", expectedCustomer, customerID)
	assert.Equal(t, *expectedCustomer, result)
	assert.NoError(t, err)
}

func TestUpdateCustomer_Error(t *testing.T) {
	mockRepo := new(MockCustomerRepo)

	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}
	customerID := uint(1)
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}
	expectedCustomer := &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		UpdatedAt:  time.Now(),
	}
	expectedError := fmt.Errorf("failed to update customer")
	mockRepo.On("UpdateCustomer", expectedCustomer, customerID).Return(nil, expectedError)
	result, err := useCase.UpdateCustomer(customer, customerID)
	mockRepo.AssertCalled(t, "UpdateCustomer", expectedCustomer, customerID)
	assert.EqualError(t, err, expectedError.Error())
	assert.NotNil(t, result)
}

func (m *MockCustomerRepo) DeleteCustomer(id uint) (interface{}, error) {
	args := m.Called(id)
	result := args.Get(0)
	err := args.Error(1)
	return result, err
}

func TestDeleteCustomer(t *testing.T) {

	mockRepo := new(MockCustomerRepo)

	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	var id uint
	id = 1

	mockRepo.On("DeleteCustomer", id).Return(nil, nil)

	result, err := useCase.DeleteCustomer(id)

	mockRepo.AssertCalled(t, "DeleteCustomer", id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteCustomer_Error(t *testing.T) {

	mockRepo := new(MockCustomerRepo)

	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	var id uint
	id = 1

	expectedError := errors.New("failed to delete customer")

	mockRepo.On("DeleteCustomer", id).Return(nil, expectedError)

	result, err := useCase.DeleteCustomer(id)

	mockRepo.AssertCalled(t, "DeleteCustomer", id)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}
