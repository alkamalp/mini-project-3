package repository

import (
	"github.com/alkamalp/crm-golang/entity"
	"gorm.io/gorm"
)

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{
		db: dbCrud,
	}
}

type ActorInterfaceRepo interface {
	CreateActor(actor *entity.Actor) (*entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	UpdateActor(actor *entity.Actor, id uint) (*entity.Actor, error)
	DeleteActor(username string) (any, error)
	LoginActor(actor *entity.Actor) (*entity.Actor, error)
}

// CreateActor new Actor
func (repo Actor) CreateActor(actor *entity.Actor) (*entity.Actor, error) {
	err := repo.db.Model(&entity.Actor{}).Create(actor).Error
	return actor, err
}

// GetActorById get single Actor by id
func (repo Actor) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	repo.db.First(&actor, "id = ? ", id)
	return actor, nil
}

// UpdateActor multiple fields
func (repo Actor) UpdateActor(actor *entity.Actor, id uint) (*entity.Actor, error) {
	err := repo.db.Model(&entity.Actor{}).Where("id = ?", id).
		Updates(actor).Error
	return nil, err
}

// DeleteActor by Id and email
func (repo Actor) DeleteActor(username string) (any, error) {
	err := repo.db.Model(&entity.Actor{}).
		Where("username = ?", username).
		Delete(&entity.Actor{}).
		Error
	return nil, err
}

func (repo Actor) LoginActor(actor *entity.Actor) (*entity.Actor, error) {
	err := repo.db.Model(&entity.Actor{}).Where("username = ? AND password = ?", actor.Username, actor.Password).Error
	return actor, err
}
