package main

import (
	"valueLabs/controllers"
	"valueLabs/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()

	r.POST("/accounts", controllers.CreateAccount)
	r.GET("/accounts/:account_id", controllers.GetAccount)
	r.POST("/transactions", controllers.CreateTransaction)

	r.Run(":8080")
}
