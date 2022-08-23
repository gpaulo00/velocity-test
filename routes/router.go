package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gpaulo00/velocity-test/app/controllers"
)

func NewRouter() *gin.Engine {
	// Register Controllers
	order := new(controllers.OrderController)

	// Declare Middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	// Declare routes
	v1 := router.Group("v1")
	{
		orderGroup := v1.Group("order")
		{
			orderGroup.POST("/process/:id", order.Process)
		}
	}
	return router

}
