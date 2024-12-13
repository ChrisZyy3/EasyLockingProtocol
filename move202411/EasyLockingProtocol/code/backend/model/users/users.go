package users

import (
	"tokenVestingProtocol/global"
	"tokenVestingProtocol/model/common"

	"github.com/gofrs/uuid"
)

type Login interface {
	CreateUser() error
	GetUserInfo() error
	GetUserId() int64
}

var _ Login = new(Users)

type Users struct {
	global.GVA_MODEL
	UUID          uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                                   // 用户UUID
	Username      string         `json:"userName" gorm:"index;comment:用户登录名"`                                                                // 用户登录名
	Password      string         `json:"-"  gorm:"comment:用户登录密码"`                                                                           // 用户登录密码
	NickName      string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                                          // 用户昵称
	HeaderImg     string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`               // 用户头像
	AuthorityId   uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                                      // 用户角色ID
	Phone         string         `json:"phone"  gorm:"comment:用户手机号"`                                                                        // 用户手机号
	Email         string         `json:"email"  gorm:"comment:用户邮箱"`                                                                         // 用户邮箱
	Enable        int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                                    //用户是否被冻结 1正常 2冻结
	OriginSetting common.JSONMap `json:"originSetting" form:"originSetting" gorm:"type:text;default:null;column:origin_setting;comment:配置;"` //配置
}

func (Users) TableName() string {
	return global.GVA_CONFIG.Mysql.Prefix + "users"
}

func (u *Users) GetUserInfo() error {
	return global.GVA_DB.Table(u.TableName()).Where("id = ?", u.ID).First(&u).Error
}

func (u *Users) CreateUser() error {
	return global.GVA_DB.Table(u.TableName()).Create(&u).Error
}

func (s *Users) GetUsername() string {
	return s.Username
}

func (s *Users) GetNickname() string {
	return s.NickName
}

func (s *Users) GetUUID() uuid.UUID {
	return s.UUID
}

func (s *Users) GetUserId() int64 {
	return s.ID
}

func (s *Users) GetAuthorityId() uint {
	return s.AuthorityId
}
