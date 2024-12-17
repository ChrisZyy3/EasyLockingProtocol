package wallet

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"tokenVestingProtocol/global"
	errX "tokenVestingProtocol/model/common/errx"
	"tokenVestingProtocol/model/common/response"
	"tokenVestingProtocol/model/wallet"
	"tokenVestingProtocol/service"

	"tokenVestingProtocol/utils"

	"github.com/gin-gonic/gin"
)

type WalletApi struct {
}

type WalletReq struct {
	Address    string `json:"address"  gorm:"column:address;type:varchar(128);comment:钱包地址" binding:"required"` // 钱包地址
	WalletName string `json:"walletName"  gorm:"column:wallet_name;type:varchar(128);comment:钱包名称"`
	IMG        string `json:"img"  gorm:"column:img;type:varchar(128);comment:钱包头像"`
}

type WalletResp struct {
	ID         int64     `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 更新时间
	Address    string    `json:"address"  gorm:"column:address;type:varchar(128);comment:钱包地址" binding:"required"` // 钱包地址
	WalletName string    `json:"walletName"  gorm:"column:wallet_name;type:varchar(128);comment:钱包名称"`
	LastSignin time.Time `json:"lastSignin"  gorm:"column:last_signin;type:varchar(128);comment:上次登录时间"`
	IMG        string    `json:"img"  gorm:"column:img;type:varchar(128);comment:钱包头像"`
}

type Output struct {
	AccessToken string      `json:"access_token"`
	Wallet      *WalletResp `json:"user_info"`
}

func (w *WalletApi) LoginApi(c *gin.Context) {
	var loginInput WalletReq
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		response.Result(400, nil, "参数验证错误", c)
		return
	}

	if err := validate(&loginInput); err != nil {
		response.Result(400, nil, err.Error(), c)
		return
	}

	output, err := LoginEnd(loginInput.Address, loginInput.WalletName, loginInput.IMG)
	if err != nil {
		response.Result(400, nil, err.Error(), c)
		return
	}
	response.Result(200, output, "登录成功", c)
}

func validate(r *WalletReq) error {
	address := strings.TrimSpace(r.Address)
	if address == "" {
		return errors.New("地址不能为空")
	}

	img := strings.TrimSpace(r.IMG)
	if img == "" {
		return errors.New("头像不能为空")

	}

	walletName := strings.TrimSpace(r.WalletName)
	if walletName == "" {
		return errors.New("钱包名称不能为空")
	}
	return nil
}

func LoginEnd(address, walletName, img string) (Output, error) {
	var Output Output
	w := wallet.Wallet{
		Address: address,
	}
	err := w.GetWalletInfoByAddress()
	if errors.Is(err, errX.ErrRecordNotFound) {
		//新用户
		w.Address = address
		w.IMG = img
		w.WalletName = walletName
		w.LastSignin = time.Now()
		tx := global.GVA_DB.Begin()
		e := tx.Create(&w).Error
		if e != nil {
			tx.Rollback()
			return Output, errors.New("存入新用户失败！" + e.Error())
		}
		//TODO:新用户其他逻辑
		tx.Commit()
	} else {
		//已存在的用户
		w.LastSignin = time.Now()
		e := global.GVA_DB.Save(&w).Error
		if e != nil {
			return Output, errors.New("用户更新登陆时间失败！" + e.Error())
		}
	}

	accessToken, _, err := utils.LoginToken(w.ID)
	if err != nil {
		fmt.Println(err.Error(), "login-end-error")
		return Output, err
	}

	Output.AccessToken = accessToken
	Output.Wallet = &WalletResp{
		ID:         w.ID,
		Address:    w.Address,
		WalletName: w.WalletName,
		IMG:        w.IMG,
		LastSignin: w.LastSignin,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt,
	}
	return Output, nil
}

var walletService = service.ServiceGroupApp.WalletServiceGroup.WalletService

func (w *WalletApi) GetWalletInfoApi(c *gin.Context) {
	walletId := utils.GetUserID(c)
	resp, err := walletService.GetWalletInfo(walletId)
	if err != nil {
		response.FailWithMessage("获取钱包信息失败", c)
		return
	}
	response.OkWithData(resp, c)

}
