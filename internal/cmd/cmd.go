package cmd

import (
	"context"
	"goshop/internal/controller"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
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
			gfUserToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			s.Group("/backend", func(group *ghttp.RouterGroup) {
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
					controller.Rotation.ListBackend, //轮播图
					controller.Rotation.Create,
					controller.Rotation.Delete,
					controller.Rotation.Update,
					controller.Position,        //手工位
					controller.Admin,           //管理员
					controller.Dashboard,       //数据大屏
					controller.Role,            //角色
					controller.Permission,      //权限
					controller.File,            //文件上传
					controller.Upload,          //文件上云
					controller.User.List,       //用户列表
					controller.User.Update,     //更新用户
					controller.User.Delete,     //删除用户
					controller.Order.List,      //订单列表
					controller.Category.List,   //分类列表
					controller.Category.Add,    //添加分类
					controller.Category.Delete, //删除分类
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
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				// 不需要frontend鉴权的路由
				group.Bind(
					controller.Rotation.ListFrontend,
					controller.Category.ListWithParentId,
				)
				// 需要frontend鉴权的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfUserToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.User.Create,        //用户注册
						controller.User.Info,          //用户信息
						controller.User.ResetPassword, //重置密码
						controller.Order.ListFrontend, //订单列表（仅用户自己的订单）
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
