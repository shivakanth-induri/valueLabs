package routes

import (
	"valueLabs/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Account routes
	router.POST("/accounts", controllers.CreateAccount)
	router.GET("/accounts/:account_id", controllers.GetAccount)
	// Transaction routes
	router.POST("/transactions", controllers.CreateTransaction)
}
