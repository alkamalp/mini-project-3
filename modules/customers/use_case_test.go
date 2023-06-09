package customers

import (
	"errors"
	"testing"
	"time"

	"github.com/alkamalp/crm-golang/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	// Set up expectations for the mock repository
	// hashedPassword := "hashed_password"
	mockRepo.On("CreateCustomer", mock.AnythingOfType("*entity.Customer")).Return(&entity.Customer{}, nil)
	// middleware.On("HashPassword", customerParam.Password).Return(hashedPassword, nil)

	// Call the function being tested
	createdCustomer, err := useCase.CreateCustomer(customer)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "CreateCustomer", mock.AnythingOfType("*entity.Customer"))

	// Assert the expected output
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
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customerID := uint(1)
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	// Set up expectations for the mock repository
	mockRepo.On("GetCustomerById", customerID).Return(customer, nil)

	// Call the function being tested
	result, err := useCase.GetCustomerById(customerID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "GetCustomerById", customerID)

	// Assert the expected output
	assert.Equal(t, customer, result)
	assert.NoError(t, err)
}

func TestGetCustomerById_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customerID := uint(1)
	expectedError := errors.New("failed to get customer")

	// Set up expectations for the mock repository
	mockRepo.On("GetCustomerById", customerID).Return(entity.Customer{}, expectedError)

	// Call the function being tested
	result, err := useCase.GetCustomerById(customerID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "GetCustomerById", customerID)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Customer{}, result)
}

func (m *MockCustomerRepo) UpdateCustomer(customer *entity.Customer, id uint) (interface{}, error) {
	args := m.Called(customer, id)
	result := args.Get(0)
	err := args.Error(1)
	return result, err
}

func TestUpdateCustomer(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customerID := uint(1)
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	// Create a new instance of the entity.Customer with the updated values
	updatedCustomer := &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Set up expectations for the mock repository
	mockRepo.On("UpdateCustomer", updatedCustomer, customerID).Return(updatedCustomer, nil)

	// Call the function being tested
	result, err := useCase.UpdateCustomer(customer, customerID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "UpdateCustomer", updatedCustomer, customerID)

	// Assert the expected output
	assert.Equal(t, updatedCustomer, result)
	assert.NoError(t, err)
}

func TestUpdateCustomer_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	customerID := uint(1)
	customer := CustomerParam{
		First_name: "John",
		Last_name:  "Doe",
		Email:      "john.doe@example.com",
		Avatar:     "avatar.jpg",
	}

	expectedError := errors.New("failed to update customer")

	// Create a new instance of the entity.Customer with the updated values
	updatedCustomer := &entity.Customer{
		First_name: customer.First_name,
		Last_name:  customer.Last_name,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
	}

	// Set up expectations for the mock repository
	mockRepo.On("UpdateCustomer", updatedCustomer, customerID).Return(nil, expectedError)

	// Call the function being tested
	result, err := useCase.UpdateCustomer(customer, customerID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "UpdateCustomer", updatedCustomer, customerID)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func (m *MockCustomerRepo) DeleteCustomer(id uint) (interface{}, error) {
	args := m.Called(id)
	result := args.Get(0)
	err := args.Error(1)
	return result, err
}

func TestDeleteCustomer(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	var id uint
	id = 1

	// Set up expectations for the mock repository
	mockRepo.On("DeleteCustomer", id).Return(nil, nil)

	// Call the function being tested
	result, err := useCase.DeleteCustomer(id)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "DeleteCustomer", id)

	// Assert the expected output
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteCustomer_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockCustomerRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseCustomer{
		customerRepo: mockRepo,
	}

	// Create test input data
	var id uint
	id = 1

	expectedError := errors.New("failed to delete customer")

	// Set up expectations for the mock repository
	mockRepo.On("DeleteCustomer", id).Return(nil, expectedError)

	// Call the function being tested
	result, err := useCase.DeleteCustomer(id)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "DeleteCustomer", id)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}
