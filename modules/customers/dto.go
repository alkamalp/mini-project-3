package customers

import (
	"github.com/alkamalp/crm-golang/dto"
	"github.com/alkamalp/crm-golang/entity"
)

type CustomerParam struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data entity.Customer `json:"data"`
}
