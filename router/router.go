package router

import (
	"gocrud/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	server         *gin.Engine
	loanController controllers.LoanControllerInterface
}

func NewRouter(server *gin.Engine, loanController controllers.LoanControllerInterface) *Router {
	return &Router{
		server,
		loanController,
	}
}

func (r *Router) Init() {
	basePath := r.server.Group("/v1")

	loans := basePath.Group("/loans")
	{
		loans.POST("/", r.loanController.Create)
	}
}
