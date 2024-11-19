package v1

import "github.com/gin-gonic/gin"

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
