package actors

import (
	"net/http"
	"strconv"

	"github.com/alkamalp/crm-golang/dto"
	"github.com/alkamalp/crm-golang/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandlerActor struct {
	ctr ControllerActor
}

func NewActorRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerActor {
	return RequestHandlerActor{
		ctr: controllerActor{
			actorUseCase: useCaseActor{
				actorRepo: repository.NewActor(dbCrud),
			},
		}}
}

func (h RequestHandlerActor) CreateActor(c *gin.Context) {
	request := ActorParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) GetActorById(c *gin.Context) {

	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetActorById(uint(actorId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) DeleteActor(c *gin.Context) {
	username := c.Param("username")
	res, err := h.ctr.DeleteActor(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) UpdateActor(c *gin.Context) {
	request := ActorParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.UpdateActor(request, uint(actorId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerActor) LoginActor(c *gin.Context) {
	request := ActorParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.LoginActor(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.Header("Authorization", res.Data)
	c.JSON(http.StatusOK, res)
}
