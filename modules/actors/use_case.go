package actors

import (
	"time"

	"github.com/alkamalp/crm-golang/entity"
	"github.com/alkamalp/crm-golang/middleware"
	"github.com/alkamalp/crm-golang/repository"
)

type UseCaseActor interface {
	CreateActor(actor ActorParam) (entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	UpdateActor(actor ActorParam, id uint) (*entity.Actor, error)
	DeleteActor(username string) (any, error)
	LoginActor(actor ActorParam) (entity.Actor, error)
}

type useCaseActor struct {
	actorRepo repository.ActorInterfaceRepo
}

func (uc useCaseActor) CreateActor(actor ActorParam) (entity.Actor, error) {
	var newActor *entity.Actor

	// Hashing password
	hashedPassword, err := middleware.HashPassword(actor.Password)
	if err != nil {
		return entity.Actor{}, err
	}

	newActor = &entity.Actor{
		Username:  actor.Username,
		Password:  hashedPassword,
		Role_id:   2,
		Verified:  0,
		Active:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = uc.actorRepo.CreateActor(newActor)
	if err != nil {
		return *newActor, err
	}
	return *newActor, nil
}

func (uc useCaseActor) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	actor, err := uc.actorRepo.GetActorById(id)
	return actor, err
}

func (uc useCaseActor) UpdateActor(actor ActorParam, id uint) (*entity.Actor, error) {
	var editActor *entity.Actor
	editActor = &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
		Role_id:  actor.Role_id,
		Verified: actor.Verified,
		Active:   actor.Active,
	}

	_, err := uc.actorRepo.UpdateActor(editActor, id)
	if err != nil {
		return editActor, err
	}
	return editActor, nil
}

func (uc useCaseActor) DeleteActor(username string) (any, error) {
	_, err := uc.actorRepo.DeleteActor(username)
	return nil, err
}

func (uc useCaseActor) LoginActor(actor ActorParam) (entity.Actor, error) {
	var newActor *entity.Actor

	newActor = &entity.Actor{
		Username: actor.Username,
		Password: actor.Password,
	}

	newActor, err := uc.actorRepo.LoginActor(newActor)
	if err != nil {
		return *newActor, err
	}
	return *newActor, nil
}
