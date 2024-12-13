package errx

import (
	"gorm.io/gorm/logger"
)

type ErrData struct {
	Msg string `json:"msg"`
	Err error  `json:"err"`
}

func (err *ErrData) Error() string {
	if err.Err != nil {
		return err.Msg + ": " + err.Err.Error()
	}
	return err.Msg
}

var (
	// 数据库相关错误码
	ErrRecordNotFound = logger.ErrRecordNotFound

	//任务
	SIGND = "已经签到过了"

	//操作
	IllegalOperation    = "非法操作"
	RepetitiveOperation = "重复操作"
	Repeatedly          = "频繁操作"
	DelErr              = "删除错误"
	NotAuthorized       = "未授权"

	//用户
	GET_USER_ERROR    = "没有此用户"
	NOT_SUFF_FUND     = "余额不足"
	USER_NO_AUTHORITY = "没有权限"

	//数据
	PutDataError = "参数错误"
	DataNotFound = "数据没找到"
	Used         = "已使用"

	//服务器
	SaveError = "保存失败"

	//订单
	NonPayment = "未支付"
)
