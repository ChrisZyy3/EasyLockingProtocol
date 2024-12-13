package vesting

import (
	v1 "tokenVestingProtocol/api/v1"
)

type RouterGroup struct {
	VestingRouter
}

var (
	vestingApi = v1.ApiGroupApp.VestingApiGroup
)
