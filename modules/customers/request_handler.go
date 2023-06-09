package customers

import (
	"net/http"
	"strconv"

	"github.com/alkamalp/crm-golang/dto"
	"github.com/alkamalp/crm-golang/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestHandlerCustomer struct {
	ctr ControllerCustomer
}

func NewCustomerRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctr: controllerCustomer{
			customerUseCase: useCaseCustomer{
				customerRepo: repository.NewCustomer(dbCrud),
			},
		}}
}

func (h RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) GetCustomerById(c *gin.Context) {
	actorId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetCustomerById(uint(actorId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) DeleteCustomer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.DeleteCustomer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) UpdateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.UpdateCustomer(request, uint(customerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
