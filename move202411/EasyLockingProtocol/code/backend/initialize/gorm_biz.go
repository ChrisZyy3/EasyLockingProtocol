package initialize

import "tokenVestingProtocol/global"

// 数据库自动迁移
func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
