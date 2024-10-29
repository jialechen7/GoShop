package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Context 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
	Admin   *ContextAdmin  // 上下文管理员信息
	Data    g.Map          // 自定KV变量，业务模块根据需要设置，不固定
}

// ContextAdmin 请求上下文中的管理员信息
type ContextAdmin struct {
	Id      int    // 管理员ID
	Name    string // 管理员账号
	RoleIds string // 管理员ids
	IsAdmin int    // 是否是管理员
}
