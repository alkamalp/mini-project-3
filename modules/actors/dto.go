package actors

import (
	"github.com/alkamalp/crm-golang/dto"
	"github.com/alkamalp/crm-golang/entity"
)

type ActorParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role_id  uint   `json:"role_id"`
	Verified int    `json:"verified"`
	Active   int    `json:"active"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data ActorParam `json:"data"`
}

type FindActor struct {
	dto.ResponseMeta
	Data entity.Actor `json:"data"`
}

type SuccessLogin struct {
	dto.ResponseMeta
	Data string `json:"data"`
}
