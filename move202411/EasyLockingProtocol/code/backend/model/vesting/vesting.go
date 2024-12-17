package vesting

import (
	"time"
	"tokenVestingProtocol/global"
)

type Vesting struct {
	global.GVA_MODEL
	TokenAddress                string    `json:"tokenAddress" form:"tokenAddress" gorm:"column:token_address;type:varchar(128);comment:token地址" binding:"required"`
	VestingDuration             int64     `json:"vestingDuration" form:"vestingDuration" gorm:"column:vesting_duration;type:bigint;comment:代币释放时间"`
	VestingDurationEndTag       int64     `json:"vestingDurationEndTag" form:"vestingDurationEndTag" gorm:"column:vesting_duration_end_tag;type:bigint;comment:代币释放时间结束标记1=second 2=minutes 3=hhour 4=days 5=weeks 6=bi-weeks 7=months 8=quarters 9=years" binding:"required"`
	UnlockSchedule              int64     `json:"unlockSchedule" form:"unlockSchedule" gorm:"column:unlock_schedule;type:bigint;comment:解锁计划1=persecond 2=perminute 3=perhour 4=perday 5=perweek 6=perbiweek 7=permonth 8=perquarter 9=peryear" binding:"required"`
	VestingEnd                  time.Time `json:"vestingEnd" form:"vestingEnd" gorm:"column:vesting_end;type:timestamp;comment:代币释放结束时间" `
	VestingStart                time.Time `json:"vestingStart" form:"vestingStart" gorm:"column:vesting_start;type:timestamp;comment:代币释放开始时间" binding:"required"`
	IsStartUponContractCreation bool      `json:"isStartUponContractCreation" form:"isStartUponContractCreation" gorm:"column:is_start_upon_contract_creation;type:tinyint;comment:是否在合约创建后立即开始释放" binding:"required"`
	IsAddCliffAmount            bool      `json:"isAddCliffAmount" form:"isAddCliffAmount" gorm:"column:is_add_cliff_amount;type:tinyint;comment:是否在 cliff 时间后添加释放金额" binding:"required"`
	CliffAmount                 float64   `json:"cliffAmount" form:"cliffAmount" gorm:"column:cliff_amount;type:decimal(18,2);comment: cliff金额" binding:"required"`
	IsAutoClaim                 bool      `json:"isAutoClaim" form:"isAutoClaim" gorm:"column:is_auto_claim;type:tinyint;comment:是否自动领取释放金额" binding:"required"`
	WhoCancel                   string    `json:"whoCancel" form:"whoCancel" gorm:"column:who_cancel;type:varchar(128);comment:谁取消合同1=唯一发件人2=唯一收件人3=两者都可以4=两者都不可以" binding:"required"`
	WhoChange                   string    `json:"whoChange" form:"whoChange" gorm:"column:who_change;type:varchar(128);comment:谁修改合同1=唯一发件人2=唯一收件人3=两者都可以4=两者都不可以" binding:"required"`
	VestingNum                  float64   `json:"num" form:"num" gorm:"column:num;type:decimal(18,2);comment:数量" binding:"required"`
	ReciverAddress              string    `json:"reciverAddress" form:"reciverAddress" gorm:"column:reciver_address;type:varchar(128);comment:收款人地址" binding:"required"`
	ContractTitle               string    `json:"contractTitle" form:"contractTitle" gorm:"column:contract_title;type:varchar(128);comment:合同标题" binding:"required"`
	ReciverEmail                string    `json:"reciverEmail" form:"reciverEmail" gorm:"column:reciver_email;type:varchar(128);comment:收款人邮箱" binding:"required"`
	WalletId                    int64     `json:"walletId" form:"walletId" gorm:"column:wallet_id;type:bigint;comment:钱包id" binding:"required"`
}

type VestingCre struct {
	TokenAddress                string    `json:"tokenAddress" form:"tokenAddress" gorm:"column:token_address;type:varchar(128);comment:token地址" binding:"required"`
	VestingDuration             int64     `json:"vestingDuration" form:"vestingDuration" gorm:"column:vesting_duration;type:bigint;comment:代币释放时间"`
	VestingDurationEndTag       int64     `json:"vestingDurationEndTag" form:"vestingDurationEndTag" gorm:"column:vesting_duration_end_tag;type:bigint;comment:代币释放时间结束标记1=second 2=minutes 3=hhour 4=days 5=weeks 6=bi-weeks 7=months 8=quarters 9=years" binding:"required"`
	UnlockSchedule              int64     `json:"unlockSchedule" form:"unlockSchedule" gorm:"column:unlock_schedule;type:bigint;comment:解锁计划1=persecond 2=perminute 3=perhour 4=perday 5=perweek 6=perbiweek 7=permonth 8=perquarter 9=peryear" binding:"required"`
	VestingEnd                  time.Time `json:"vestingEnd" form:"vestingEnd" gorm:"column:vesting_end;type:timestamp;comment:代币释放结束时间" `
	VestingStart                time.Time `json:"vestingStart" form:"vestingStart" gorm:"column:vesting_start;type:timestamp;comment:代币释放开始时间" binding:"required"`
	IsStartUponContractCreation bool      `json:"isStartUponContractCreation" form:"isStartUponContractCreation" gorm:"column:is_start_upon_contract_creation;type:tinyint;comment:是否在合约创建后立即开始释放" binding:"required"`
	IsAddCliffAmount            bool      `json:"isAddCliffAmount" form:"isAddCliffAmount" gorm:"column:is_add_cliff_amount;type:tinyint;comment:是否在 cliff 时间后添加释放金额" binding:"required"`
	CliffAmount                 float64   `json:"cliffAmount" form:"cliffAmount" gorm:"column:cliff_amount;type:decimal(18,2);comment: cliff金额" binding:"required"`
	IsAutoClaim                 bool      `json:"isAutoClaim" form:"isAutoClaim" gorm:"column:is_auto_claim;type:tinyint;comment:是否自动领取释放金额" binding:"required"`
	WhoCancel                   string    `json:"whoCancel" form:"whoCancel" gorm:"column:who_cancel;type:varchar(128);comment:谁取消合同1=唯一发件人2=唯一收件人3=两者都可以4=两者都不可以" binding:"required"`
	WhoChange                   string    `json:"whoChange" form:"whoChange" gorm:"column:who_change;type:varchar(128);comment:谁修改合同1=唯一发件人2=唯一收件人3=两者都可以4=两者都不可以" binding:"required"`
	VestingNum                  float64   `json:"num" form:"num" gorm:"column:num;type:decimal(18,2);comment:数量" binding:"required"`
	ReciverAddress              string    `json:"reciverAddress" form:"reciverAddress" gorm:"column:reciver_address;type:varchar(128);comment:收款人地址" binding:"required"`
	ContractTitle               string    `json:"contractTitle" form:"contractTitle" gorm:"column:contract_title;type:varchar(128);comment:合同标题" binding:"required"`
	ReciverEmail                string    `json:"reciverEmail" form:"reciverEmail" gorm:"column:reciver_email;type:varchar(128);comment:收款人邮箱" binding:"required"`
	WalletId                    int64     `json:"walletId" form:"walletId" gorm:"column:wallet_id;type:bigint;comment:钱包id" binding:"required"`
}

func (Vesting) TableName() string {
	return global.GVA_CONFIG.Mysql.Prefix + "vesting"
}

func (v *Vesting) GetVestingById() error {
	return global.GVA_DB.Table(v.TableName()).Where("id = ?", v.ID).First(&v).Error
}

func (v *Vesting) GetVestingList(page int, pageSize int) ([]Vesting, error) {
	data := make([]Vesting, 0)
	if err := global.GVA_DB.Table(v.TableName()).Limit(pageSize).Offset((page - 1) * pageSize).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (v *Vesting) CreateVesting() error {
	return global.GVA_DB.Table(v.TableName()).Create(&v).Error
}

func (v *Vesting) ToArgs(v1 VestingCre) {
	v.TokenAddress = v1.TokenAddress
	v.VestingDuration = v1.VestingDuration
	v.VestingDurationEndTag = v1.VestingDurationEndTag
	v.UnlockSchedule = v1.UnlockSchedule
	v.VestingEnd = v1.VestingEnd
	v.VestingStart = v1.VestingStart
	v.IsStartUponContractCreation = v1.IsStartUponContractCreation
	v.IsAddCliffAmount = v1.IsAddCliffAmount
	v.CliffAmount = v1.CliffAmount
	v.IsAutoClaim = v1.IsAutoClaim
	v.WhoCancel = v1.WhoCancel
	v.WhoChange = v1.WhoChange
	v.VestingNum = v1.VestingNum
	v.ReciverAddress = v1.ReciverAddress
	v.ContractTitle = v1.ContractTitle
	v.ReciverEmail = v1.ReciverEmail
	v.WalletId = v1.WalletId
}
