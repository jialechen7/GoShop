package cmd

import (
	"context"
	"goshop/internal/controller"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goshop/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().CORS,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					hello.NewV1(),           //示例
					controller.Rotation,     //轮播图
					controller.Position,     //手工位
					controller.Admin.Create, //管理员创建
					controller.Admin.List,   //管理员列表
					controller.Login,        //登录
				)
				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/backend/admin/info":   controller.Admin.Info,
						"/backend/admin/update": controller.Admin.Update,
						"/backend/admin/delete": controller.Admin.Delete,
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
