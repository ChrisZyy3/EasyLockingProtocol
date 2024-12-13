package wallet

import (
	"time"
	"tokenVestingProtocol/global"
)

type Login interface {
}

var _ Login = new(Wallet)

type Wallet struct {
	global.GVA_MODEL
	Address    string    `json:"address"  gorm:"column:address;type:varchar(128);comment:钱包地址" binding:"required"` // 钱包地址
	WalletName string    `json:"wallet_name"  gorm:"column:wallet_name;type:varchar(128);comment:钱包名称"`
	LastSignin time.Time `json:"last_signin"  gorm:"column:last_signin;type:timestamp;comment:上次登录时间"`
	IMG        string    `json:"img"  gorm:"column:img;type:varchar(128);comment:钱包头像"`
}

func (Wallet) TableName() string {
	return global.GVA_CONFIG.Mysql.Prefix + "wallet"
}

func (w *Wallet) GetWalletInfoById() error {
	return global.GVA_DB.Table(w.TableName()).Where("id = ?", w.ID).First(&w).Error
}

func (w *Wallet) GetWalletInfoByAddress() error {
	return global.GVA_DB.Table(w.TableName()).Where("address = ?", w.Address).First(&w).Error
}

func (w *Wallet) CreateWallet() error {
	return global.GVA_DB.Table(w.TableName()).Create(&w).Error
}
