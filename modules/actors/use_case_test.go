package actors

import (
	"errors"
	"testing"
	"time"

	"github.com/alkamalp/crm-golang/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


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
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actorParam := ActorParam{
		Username: "testuser",
		Password: "password",
	}

	
	
	mockRepo.On("CreateActor", mock.AnythingOfType("*entity.Actor")).Return(&entity.Actor{}, nil)
	

	
	createdActor, err := useCase.CreateActor(actorParam)

	
	mockRepo.AssertCalled(t, "CreateActor", mock.AnythingOfType("*entity.Actor"))

	
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
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actorID := uint(1)
	actor := entity.Actor{
		ID:        actorID,
		Username:  "John Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	
	mockRepo.On("GetActorById", actorID).Return(actor, nil)

	
	result, err := useCase.GetActorById(actorID)

	
	mockRepo.AssertCalled(t, "GetActorById", actorID)

	
	assert.Equal(t, actor, result)
	assert.NoError(t, err)
}

func TestGetActorById_Error(t *testing.T) {
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actorID := uint(1)
	expectedError := errors.New("failed to get actor")

	
	mockRepo.On("GetActorById", actorID).Return(entity.Actor{}, expectedError)

	
	result, err := useCase.GetActorById(actorID)

	
	mockRepo.AssertCalled(t, "GetActorById", actorID)

	
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
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actorID := uint(1)
	actor := ActorParam{
		Username: "JohnDoe",
		Password: "password",
		Role_id:  2,
		Verified: 0,
		Active:   0,
	}

	
	updatedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
		Role_id:  actor.Role_id,
		Verified: actor.Verified,
		Active:   actor.Active,
	}

	
	mockRepo.On("UpdateActor", updatedActor, actorID).Return(updatedActor, nil)

	
	result, err := useCase.UpdateActor(actor, actorID)

	
	mockRepo.AssertCalled(t, "UpdateActor", updatedActor, actorID)

	
	assert.Equal(t, updatedActor, result)
	assert.NoError(t, err)
}

func TestUpdateActor_Error(t *testing.T) {
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actorID := uint(1)
	actor := ActorParam{
		Username: "JohnDoe",
		Password: "password",
		Role_id:  2,
		Verified: 0,
		Active:   0,
	}

	expectedError := errors.New("failed to update actor")

	
	updatedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
		Role_id:  actor.Role_id,
		Verified: actor.Verified,
		Active:   actor.Active,
	}

	
	mockRepo.On("UpdateActor", updatedActor, actorID).Return(nil, expectedError)

	
	result, err := useCase.UpdateActor(actor, actorID)

	
	mockRepo.AssertCalled(t, "UpdateActor", updatedActor, actorID)

	
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
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	username := "JohnDoe"

	
	mockRepo.On("DeleteActor", username).Return(nil, nil)

	
	result, err := useCase.DeleteActor(username)

	
	mockRepo.AssertCalled(t, "DeleteActor", username)

	
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteActor_Error(t *testing.T) {
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	username := "JohnDoe"

	expectedError := errors.New("failed to delete actor")

	
	mockRepo.On("DeleteActor", username).Return(nil, expectedError)

	
	result, err := useCase.DeleteActor(username)

	
	mockRepo.AssertCalled(t, "DeleteActor", username)

	
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
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actor := ActorParam{
		Username: "john",
		Password: "password",
	}

	
	expectedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
	}

	
	mockRepo.On("LoginActor", expectedActor).Return(expectedActor, nil)

	
	result, err := useCase.LoginActor(actor)

	
	mockRepo.AssertCalled(t, "LoginActor", expectedActor)

	
	assert.NoError(t, err)
	assert.Equal(t, *expectedActor, result)
}

func TestLoginActor_Error(t *testing.T) {
	
	mockRepo := new(MockActorRepo)

	
	useCase := useCaseActor{
		actorRepo: mockRepo,
	}

	
	actor := ActorParam{
		Username: "john",
		Password: "password",
	}

	
	expectedActor := &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
	}

	expectedError := errors.New("login failed")

	
	mockRepo.On("LoginActor", expectedActor).Return(nil, expectedError)

	
	result, err := useCase.LoginActor(actor)

	
	mockRepo.AssertCalled(t, "LoginActor", expectedActor)

	
	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, entity.Actor{}, result)
}
