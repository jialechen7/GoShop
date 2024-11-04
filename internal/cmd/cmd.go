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
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// gtoken 中间件绑定
				err := gfAdminToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					hello.NewV1(),         //示例
					controller.Rotation,   //轮播图
					controller.Position,   //手工位
					controller.Admin,      //管理员
					controller.Dashboard,  //数据大屏
					controller.Role,       //角色
					controller.Permission, //权限
					controller.File,       //文件上传
					//controller.Login,        //登录（使用gtoken时不需要绑定，在gtoken中绑定）
				)

				//// Special handler that needs authentication.
				//group.Group("/", func(group *ghttp.RouterGroup) {
				//	//group.Middleware(service.Middleware().Auth)  // 使用gtoken时，不需要绑定jwt中间件
				//	group.ALLMap(g.Map{
				//		"/backend/admin/info":   controller.Admin.Info,
				//		"/backend/admin/update": controller.Admin.Update,
				//		"/backend/admin/delete": controller.Admin.Delete,
				//	})
				//})
			})
			s.Run()
			return nil
		},
	}
)
