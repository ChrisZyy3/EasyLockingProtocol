package v1

import (
	"tokenVestingProtocol/api/v1/vesting"
	"tokenVestingProtocol/api/v1/wallet"
)

type ApiGroup struct {
	WalletApiGroup  wallet.ApiGroup
	VestingApiGroup vesting.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
