package wallet

import (
	"tokenVestingProtocol/global"
	"tokenVestingProtocol/model/wallet"
)

type WalletService struct{}

var WalletServiceApp = new(WalletService)

func (w *WalletService) GetWalletInfo(walletId int64) (*wallet.Wallet, error) {
	wallet := &wallet.Wallet{
		GVA_MODEL: global.GVA_MODEL{ID: walletId},
	}
	if err := wallet.GetWalletInfoById(); err != nil {
		return nil, err
	}
	return wallet, nil
}

func (w *WalletService) CreateWallet(Wallet *wallet.Wallet) error {
	if err := Wallet.CreateWallet(); err != nil {
		return err
	}
	return nil
}

func (w *WalletService) FindWalletById(id int) (*wallet.Wallet, error) {
	var Wallet wallet.Wallet
	if err := global.GVA_DB.Where("id = ?", id).First(&Wallet).Error; err != nil {
		return nil, err
	}
	return &Wallet, nil
}
