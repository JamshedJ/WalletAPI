package v1

import (
	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/gin-gonic/gin"
)

func (ctrl *ControllerV1) CheckWalletExists(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "userID is required"})
		return
	}

	isWalletExists, err := ctrl.Services.Wallet.CheckWalletExists(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"is_wallet_exists": isWalletExists})
}

func (ctrl *ControllerV1) GetWalletBalance(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "userID is required"})
		return
	}

	wallet, err := ctrl.Services.Wallet.GetWalletBalance(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"wallet_balance": wallet.Balance})
}

func (ctrl *ControllerV1) TopUpWallet(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "userID is required"})
		return
	}

	var req = struct {
		WalletID uint    `json:"wallet_id" binding:"required"`
		Amount   float64 `json:"amount" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.Services.Wallet.TopUpWallet(c.Request.Context(), userID, &dto.TopUpWalletIn{
		Amount:   req.Amount,
		WalletID: req.WalletID,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "wallet topped up successfully"})
}

func (ctrl *ControllerV1) GetMonthlySummary(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "userID is required"})
		return
	}

	summary, err := ctrl.Services.Wallet.GetMonthlySummary(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"summary": summary})
}
