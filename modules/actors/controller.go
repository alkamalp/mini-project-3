package actors

import (
	"time"

	"github.com/alkamalp/crm-golang/dto"
	"github.com/golang-jwt/jwt"
)

type ControllerActor interface {
	CreateActor(req ActorParam) (any, error)
	GetActorById(id uint) (FindActor, error)
	UpdateActor(req ActorParam, id uint) (any, error)
	DeleteActor(username string) (any, error)
	LoginActor(req ActorParam) (SuccessLogin, error)
}

type controllerActor struct {
	actorUseCase UseCaseActor
}

func (uc controllerActor) CreateActor(req ActorParam) (any, error) {

	actor, err := uc.actorUseCase.CreateActor(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create actor",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: ActorParam{
			Username: actor.Username,
			Password: actor.Password,
			Role_id:  actor.Role_id,
			Verified: actor.Verified,
			Active:   actor.Active,
		},
	}
	return res, nil
}

func (uc controllerActor) GetActorById(id uint) (FindActor, error) {
	var res FindActor
	actor, err := uc.actorUseCase.GetActorById(id)
	if err != nil {
		return FindActor{}, err
	}
	res.Data = actor
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get actor",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerActor) UpdateActor(req ActorParam, id uint) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.actorUseCase.UpdateActor(req, id)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success update"
	res.MessageTitle = "update"

	return res, nil
}

func (uc controllerActor) DeleteActor(email string) (any, error) {
	var res dto.ResponseMeta
	_, err := uc.actorUseCase.DeleteActor(email)
	if err != nil {
		return dto.ResponseMeta{}, err
	}
	res.Success = true
	res.Message = "Success Delete"
	res.MessageTitle = "Delete"

	return res, nil
}

func (uc controllerActor) LoginActor(req ActorParam) (SuccessLogin, error) {

	actor, err := uc.actorUseCase.LoginActor(req)
	if err != nil {
		return SuccessLogin{}, err
	}
	 // Inisialisasi klaim-klaim yang ingin Anda sertakan dalam token
	 claims := jwt.MapClaims{
        "sub": actor.Role_id,
        "name": actor.Username,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(time.Hour * 1).Unix(),
    }

    // Tandatangani token dengan kunci rahasia
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte("secret-key"))
    if err != nil {
        // Penanganan kesalahan
    }

    // Gunakan signedToken seperti yang Anda butuhkan

	res := SuccessLogin{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success login actor",
			Message:      "Success actor",
			ResponseTime: "",
		},
		Data: signedToken,
	}
	return res, nil
}
