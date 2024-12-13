package vesting

import (
	v1 "tokenVestingProtocol/api/v1"

	"github.com/gin-gonic/gin"
)

type VestingRouter struct{}

func (v *VestingRouter) InitVestingRouter(Router *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {
	vestingRouter := Router.Group("vesting")
	api := v1.ApiGroupApp.VestingApiGroup.VestingApi
	{
		vestingRouter.POST("create/v1", api.CreateApi)
		vestingRouter.GET("get/v1", api.GetApi)
	}
}
