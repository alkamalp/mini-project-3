package customers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteCustomer struct {
	CustomerRequestHandeler RequestHandlerCustomer
}

func NewRouter(
	dbCrud *gorm.DB,
) RouteCustomer {
	return RouteCustomer{CustomerRequestHandeler: NewCustomerRequestHandler(
		dbCrud,
	)}
}

func (r RouteCustomer) Handle(routeVersion *gin.Engine) {
	basepath := "/customer"
	customer := routeVersion.Group(basepath)

	customer.POST("",
		r.CustomerRequestHandeler.CreateCustomer,
	)

	customer.GET("/:id",
		r.CustomerRequestHandeler.GetCustomerById,
	)
	customer.PUT("/:id",
		r.CustomerRequestHandeler.UpdateCustomer,
	)
	customer.DELETE("/:id",
		r.CustomerRequestHandeler.DeleteCustomer,
	)
}
