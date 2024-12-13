package service

import (
	"tokenVestingProtocol/service/system"
	"tokenVestingProtocol/service/users"
	"tokenVestingProtocol/service/wallet"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   users.UserServiceGroup
	WalletServiceGroup wallet.WalletServiceGroup
}
