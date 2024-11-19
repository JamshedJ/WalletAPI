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

	wallet := v1.Group("/wallet")
	{
		wallet.POST("/:id/balance", ctrl.GetWalletBalance)
		wallet.POST("/:id/exists", ctrl.CheckWalletExists)
		wallet.POST("/:id/topup", ctrl.TopUpWallet)
		wallet.POST("/:id/summary", ctrl.GetMonthlySummary)
	}

	return nil
}
