package v1

import (
	"github.com/JamshedJ/WalletAPI/domain/services"
	"github.com/gin-gonic/gin"
)

type ControllerV1 struct {
	Services *services.ServiceFacade
}

func InitRoutes(e *gin.Engine, svc *services.ServiceFacade) error {
	ctrl := &ControllerV1{
		Services: svc,
	}

	v1 := e.Group("/v1")

	wallet := v1.Group("/wallet", ctrl.AuthMiddleware())
	{
		wallet.POST("/balance", ctrl.GetWalletBalance)
		wallet.POST("/exists", ctrl.CheckWalletExists)
		wallet.POST("/topup", ctrl.TopUpWallet)
		wallet.POST("/summary", ctrl.GetMonthlySummary)
	}

	return nil
}
