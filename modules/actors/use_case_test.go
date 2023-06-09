package actors

import (
	"errors"
	"testing"
	"time"

	"github.com/alkamalp/crm-golang/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockActorRepo is a mock implementation of the ActorRepository interface
type MockActorRepo struct {
	mock.Mock
}

func (m *MockActorRepo) CreateActor(actor *entity.Actor) (*entity.Actor, error) {
	args := m.Called(actor)
	result := args.Get(0)
	err := args.Error(1)
	if result == nil {
		return nil, err
	}
	return result.(*entity.Actor), err
}

func TestCreateActor(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actorParam := ActorParam{
		Username: "testuser",
		Password: "password",
	}

	// Set up expectations for the mock repository
	// hashedPassword := "hashed_password"
	mockRepo.On("CreateActor", mock.AnythingOfType("*entity.Actor")).Return(&entity.Actor{}, nil)
	// middleware.On("HashPassword", actorParam.Password).Return(hashedPassword, nil)

	// Call the function being tested
	createdActor, err := useCase.CreateActor(actorParam)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "CreateActor", mock.AnythingOfType("*entity.Actor"))

	// Assert the expected output
	assert.NotNil(t, createdActor)
	assert.NoError(t, err)
}

func (m *MockActorRepo) GetActorById(id uint) (entity.Actor, error) {
	args := m.Called(id)
	result := args.Get(0)
	err := args.Error(1)
	if result == nil {
		return entity.Actor{}, err
	}
	return result.(entity.Actor), err
}

func TestGetActorById(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actorID := uint(1)
	actor := entity.Actor{
		ID:        actorID,
		Username:  "John Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set up expectations for the mock repository
	mockRepo.On("GetActorById", actorID).Return(actor, nil)

	// Call the function being tested
	result, err := useCase.GetActorById(actorID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "GetActorById", actorID)

	// Assert the expected output
	assert.Equal(t, actor, result)
	assert.NoError(t, err)
}

func TestGetActorById_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actorID := uint(1)
	expectedError := errors.New("failed to get actor")

	// Set up expectations for the mock repository
	mockRepo.On("GetActorById", actorID).Return(entity.Actor{}, expectedError)

	// Call the function being tested
	result, err := useCase.GetActorById(actorID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "GetActorById", actorID)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Actor{}, result)
}

func (m *MockActorRepo) UpdateActor(actor *entity.Actor, id uint) (interface{}, error) {
	args := m.Called(actor, id)
	result := args.Get(0)
	err := args.Error(1)
	return result, err
}

func TestUpdateActor(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actorID := uint(1)
	actor := ActorParam{
		Username: "JohnDoe",
		Password: "password",
		Role_id:  2,
		Verified: 0,
		Active:   0,
	}

	// Create a new instance of the entity.Actor with the updated values
	updatedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
		Role_id:  actor.Role_id,
		Verified: actor.Verified,
		Active:   actor.Active,
	}

	// Set up expectations for the mock repository
	mockRepo.On("UpdateActor", updatedActor, actorID).Return(updatedActor, nil)

	// Call the function being tested
	result, err := useCase.UpdateActor(actor, actorID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "UpdateActor", updatedActor, actorID)

	// Assert the expected output
	assert.Equal(t, updatedActor, result)
	assert.NoError(t, err)
}

func TestUpdateActor_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actorID := uint(1)
	actor := ActorParam{
		Username: "JohnDoe",
		Password: "password",
		Role_id:  2,
		Verified: 0,
		Active:   0,
	}

	expectedError := errors.New("failed to update actor")

	// Create a new instance of the entity.Actor with the updated values
	updatedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
		Role_id:  actor.Role_id,
		Verified: actor.Verified,
		Active:   actor.Active,
	}

	// Set up expectations for the mock repository
	mockRepo.On("UpdateActor", updatedActor, actorID).Return(nil, expectedError)

	// Call the function being tested
	result, err := useCase.UpdateActor(actor, actorID)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "UpdateActor", updatedActor, actorID)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func (m *MockActorRepo) DeleteActor(username string) (interface{}, error) {
	args := m.Called(username)
	result := args.Get(0)
	err := args.Error(1)
	return result, err
}

func TestDeleteActor(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	username := "JohnDoe"

	// Set up expectations for the mock repository
	mockRepo.On("DeleteActor", username).Return(nil, nil)

	// Call the function being tested
	result, err := useCase.DeleteActor(username)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "DeleteActor", username)

	// Assert the expected output
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteActor_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	username := "JohnDoe"

	expectedError := errors.New("failed to delete actor")

	// Set up expectations for the mock repository
	mockRepo.On("DeleteActor", username).Return(nil, expectedError)

	// Call the function being tested
	result, err := useCase.DeleteActor(username)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "DeleteActor", username)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func (m *MockActorRepo) LoginActor(actor *entity.Actor) (*entity.Actor, error) {
	args := m.Called(actor)
	result := args.Get(0)
	err := args.Error(1)
	return result.(*entity.Actor), err
}

func TestLoginActor(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actor := ActorParam{
		Username: "john",
		Password: "password",
	}

	// Create a new instance of the expected actor returned by the repository
	expectedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
	}

	// Set up expectations for the mock repository
	mockRepo.On("LoginActor", expectedActor).Return(expectedActor, nil)

	// Call the function being tested
	result, err := useCase.LoginActor(actor)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "LoginActor", expectedActor)

	// Assert the expected output
	assert.NoError(t, err)
	assert.Equal(t, *expectedActor, result)
}

func TestLoginActor_Error(t *testing.T) {
	// Create a new instance of the mock repository
	mockRepo := new(MockActorRepo)

	// Create an instance of the use case with the mock repository
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	// Create test input data
	actor := ActorParam{
		Username: "john",
		Password: "password",
	}

	// Create a new instance of the expected actor returned by the repository
	expectedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
	}

	expectedError := errors.New("login failed")

	// Set up expectations for the mock repository
	mockRepo.On("LoginActor", expectedActor).Return(nil, expectedError)

	// Call the function being tested
	result, err := useCase.LoginActor(actor)

	// Assert that the expected repository method was called with the correct input
	mockRepo.AssertCalled(t, "LoginActor", expectedActor)

	// Assert the expected error
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Actor{}, result)
}
