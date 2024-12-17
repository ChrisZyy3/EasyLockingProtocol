package vesting

import (
	"strconv"
	"tokenVestingProtocol/model/common/response"
	"tokenVestingProtocol/model/vesting"
	"tokenVestingProtocol/utils"

	"github.com/gin-gonic/gin"
)

type VestingApi struct {
}

func (v *VestingApi) CreateApi(c *gin.Context) {
	var req vesting.VestingCre
	var arg vesting.Vesting
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	walletId := utils.GetUserID(c)
	req.WalletId = walletId
	arg.ToArgs(req)
	if err := arg.CreateVesting(); err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData("创建成功", c)
}

func (v *VestingApi) GetApi(c *gin.Context) {
	page := c.Query("page")
	pageSize := c.Query("pageSize")

	vesting := &vesting.Vesting{}
	offsert, _ := strconv.Atoi(page)
	limit, _ := strconv.Atoi(pageSize)
	resp, err := vesting.GetVestingList(offsert, limit)
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(resp, c)
}
