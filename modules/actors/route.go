package actors

import (
	"github.com/alkamalp/crm-golang/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteActor struct {
	ActorRequestHandeler RequestHandlerActor
}

func NewRouter(
	dbCrud *gorm.DB,
) RouteActor {
	return RouteActor{ActorRequestHandeler: NewActorRequestHandler(
		dbCrud,
	)}
}

func (r RouteActor) Handle(routeVersion *gin.Engine) {
	basepath := "/actor"
	actor := routeVersion.Group(basepath)

	actor.POST("",
		r.ActorRequestHandeler.CreateActor,
	)

	actor.GET("/:id", middleware.Auth,
		r.ActorRequestHandeler.GetActorById,
	)
	actor.PUT("/:id", middleware.Auth,
		r.ActorRequestHandeler.UpdateActor,
	)
	actor.DELETE("/:username", middleware.Auth,
		r.ActorRequestHandeler.DeleteActor,
	)
	actor.POST("/login",
		r.ActorRequestHandeler.LoginActor,
	)
}
