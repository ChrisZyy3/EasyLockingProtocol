//go:build !windows
// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// initServer 初始化一个服务器实例。
// 参数:
//
//	address: 服务器监听的地址。
//	router: 用于处理HTTP请求的路由器指针。
//
// 返回值:
//
//	*server: 初始化后的服务器实例指针。
func initServer(address string, router *gin.Engine) server {
	// 创建一个endless服务器实例，允许优雅地启动和关闭。
	s := endless.NewServer(address, router)

	// 设置读取请求头的超时时间，防止慢速攻击。
	s.ReadHeaderTimeout = 10 * time.Minute
	// 设置写入响应的超时时间，确保及时响应。
	s.WriteTimeout = 10 * time.Minute
	// 设置请求头的最大字节数，防止DoS攻击。
	s.MaxHeaderBytes = 1 << 20

	// 返回初始化后的服务器实例。
	return s
}
