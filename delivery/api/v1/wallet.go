package v1

import (
	"errors"
	"strconv"

	"github.com/JamshedJ/WalletAPI/domain/dto"
	"github.com/gin-gonic/gin"
)

func (ctrl *ControllerV1) GetWalletBalance(c *gin.Context) {
	userIDStr := c.Param("id")
	if userIDStr == "" {
		handleError(c, errors.New("userID is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handleError(c, errors.New("invalid userID"))
		return
	}

	wallet, err := ctrl.Services.Wallet.GetWalletBalance(c.Request.Context(), uint(userID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"wallet_balance": wallet.Balance})
}

func (ctrl *ControllerV1) CheckWalletExists(c *gin.Context) {
	userIDStr := c.Param("id")
	if userIDStr == "" {
		handleError(c, errors.New("userID is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handleError(c, errors.New("invalid userID"))
		return
	}

	isWalletExists, err := ctrl.Services.Wallet.CheckWalletExists(c.Request.Context(), uint(userID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"is_wallet_exists": isWalletExists})
}

func (ctrl *ControllerV1) TopUpWallet(c *gin.Context) {
	userIDStr := c.Param("id")
	if userIDStr == "" {
		handleError(c, errors.New("userID is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handleError(c, errors.New("invalid userID"))
		return
	}

	var req = struct {
		WalletID uint    `json:"wallet_id" binding:"required"`
		UserID   uint    `json:"user_id" binding:"required"`
		Amount   float64 `json:"amount" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, err)
		return
	}

	err = ctrl.Services.Wallet.TopUpWallet(c.Request.Context(), uint(userID), &dto.TopUpWalletIn{
		WalletID: req.WalletID,
		UserID:   req.UserID,
		Amount:   req.Amount,
	})
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "wallet topped up successfully"})
}

func (ctrl *ControllerV1) GetMonthlySummary(c *gin.Context) {
	userIDStr := c.Param("id")
	if userIDStr == "" {
		handleError(c, errors.New("userID is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		handleError(c, errors.New("invalid userID"))
		return
	}

	summary, err := ctrl.Services.Wallet.GetMonthlySummary(c.Request.Context(), uint(userID))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, gin.H{"summary": summary})
}
