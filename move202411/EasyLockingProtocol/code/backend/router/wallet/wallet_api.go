package wallet

import (
	v1 "tokenVestingProtocol/api/v1"

	"github.com/gin-gonic/gin"
)

type WalletRouter struct {
}

func (w *WalletRouter) InitWalletRouter(PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {
	userRouter := PublicGroup.Group("wallet")
	userPrivateRouter := PrivateGroup.Group("wallet")

	api := v1.ApiGroupApp.WalletApiGroup.WalletApi
	{
		userRouter.POST("login", api.LoginApi)
		userPrivateRouter.GET("", api.GetWalletInfoApi)
	}
}
