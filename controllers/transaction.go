package controllers

import (
	"fmt"
	"net/http"
	"valueLabs/database"
	"valueLabs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransaction(c *gin.Context) {
	var req models.Transaction

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	if err := validateTransactionInput(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var src, dst models.Account
		if err := tx.First(&src, "account_id = ?", req.SourceAccountID).Error; err != nil {
			return fmt.Errorf("source account (%d) not found", req.SourceAccountID)
		}
		if err := tx.First(&dst, "account_id = ?", req.DestinationAccountID).Error; err != nil {
			return fmt.Errorf("destination account (%d) not found", req.DestinationAccountID)
		}
		if src.InitialBalance < req.Amount {
			return fmt.Errorf("insufficient funds in source account")
		}

		src.InitialBalance -= req.Amount
		dst.InitialBalance += req.Amount

		if err := tx.Save(&src).Error; err != nil {
			return fmt.Errorf("failed to update source account: %v", err)
		}
		if err := tx.Save(&dst).Error; err != nil {
			return fmt.Errorf("failed to update destination account: %v", err)
		}
		if err := tx.Create(&req).Error; err != nil {
			return fmt.Errorf("failed to record transaction: %v", err)
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction failed: " + err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func validateTransactionInput(t models.Transaction) error {
	if t.SourceAccountID <= 0 {
		return fmt.Errorf("source_account_id must be a positive number")
	}
	if t.DestinationAccountID <= 0 {
		return fmt.Errorf("destination_account_id must be a positive number")
	}
	if t.SourceAccountID == t.DestinationAccountID {
		return fmt.Errorf("source and destination account IDs cannot be the same")
	}
	if t.Amount <= 0 {
		return fmt.Errorf("amount must be a positive value")
	}
	return nil
}
