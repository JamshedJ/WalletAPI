package api

import (
	"fmt"

	v1 "github.com/JamshedJ/WalletAPI/delivery/api/v1"
	"github.com/JamshedJ/WalletAPI/domain/services"
	"github.com/gin-gonic/gin"
)

func Run(svc *services.ServiceFacade, port int) error {
	e := gin.Default()

	v1.InitRoutes(e, svc)

	return e.Run(fmt.Sprintf(":%v", port))
}
