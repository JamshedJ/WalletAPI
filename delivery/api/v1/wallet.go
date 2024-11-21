package v1

import (
	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ctrl *ControllerV1) GetWalletBalance(c *gin.Context) {
	var req struct {
		Account string `json:"account" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	wallet, err := ctrl.Services.Wallet.GetWalletBalance(c.Request.Context(), req.Account)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"wallet_balance": wallet.Balance})
}

func (ctrl *ControllerV1) CheckWalletExists(c *gin.Context) {
	var req struct {
		Account string `json:"account" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	exists, err := ctrl.Services.Wallet.CheckWalletExists(c.Request.Context(), req.Account)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"exists": exists})
}

func (ctrl *ControllerV1) TopUpWallet(c *gin.Context) {
	partnerIDStr := c.GetString("partnerID")
	partrerID, err := uuid.Parse(partnerIDStr)
	if err != nil {
		handleError(c, err)
		return
	}

	var req struct {
		Account string  `json:"account" binding:"required"`
		Amount  float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	err = ctrl.Services.Wallet.TopUpWallet(c.Request.Context(), &dto.TopUpWalletIn{
		PartnerID: partrerID,
		Account:   req.Account,
		Amount:    req.Amount,
	})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "wallet topped up successfully"})
}

func (ctrl *ControllerV1) GetMonthlySummary(c *gin.Context) {
	partnerIDStr := c.GetString("partnerID")
	partnerID, err := uuid.Parse(partnerIDStr)
	if err != nil {
		handleError(c, err)
		return
	}

	summary, err := ctrl.Services.Wallet.GetMonthlySummary(c.Request.Context(), partnerID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, summary)
}
