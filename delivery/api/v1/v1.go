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

	_ = ctrl

	v1 := e.Group("/v1")

	w := v1.Group("/wallet")
	{
		_ = w
	}
	
	return nil
}
