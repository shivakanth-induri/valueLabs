package controllers

import (
	"fmt"
	"net/http"
	"valueLabs/database"
	"valueLabs/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var acc models.Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}
	if err := validateAccountInput(acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&acc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account creation failed: " + err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func validateAccountInput(acc models.Account) error {
	if acc.AccountID <= 0 {
		return fmt.Errorf("account_id must be a positive number")
	}
	if acc.InitialBalance < 0 {
		return fmt.Errorf("initial_balance cannot be negative")
	}
	return nil
}

func GetAccount(c *gin.Context) {
	var acc models.Account
	id := c.Param("account_id")
	if err := database.DB.First(&acc, "account_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"account_id": acc.AccountID,
		"balance":    acc.InitialBalance,
	})
}
