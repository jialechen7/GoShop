package cmd

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/controller"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
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
					controller.Position,              //手工位
					controller.Admin,                 //管理员
					controller.Dashboard,             //数据大屏
					controller.Role,                  //角色
					controller.Permission,            //权限
					controller.File,                  //文件上传
					controller.Upload,                //文件上云
					controller.User.List,             //用户列表
					controller.User.Update,           //更新用户
					controller.User.Delete,           //删除用户
					controller.Order.List,            //订单列表
					controller.Category.ListAll,      //全部一级分类列表
					controller.Category.List,         //分类列表
					controller.Category.Add,          //添加分类
					controller.Category.Delete,       //删除分类
					controller.Article.ListBackend,   //文章列表
					controller.Article.UpdateBackend, //更新文章
					controller.Article.AddBackend,    //添加文章
					controller.Article.DeleteBackend, //删除文章
					controller.Comment.ListBackend,   //评论列表
					controller.Comment.DeleteBackend, //删除评论
					controller.Consignee.ListBackend, //查询收货人列表
					controller.Goods.ListBackend,     //商品列表
					controller.Goods.AddBackend,      //添加商品
					controller.Goods.UpdateBackend,   //更新商品
					controller.Goods.DeleteBackend,   //删除商品
					controller.Goods.DetailBackend,   //商品详情
					controller.Coupon.List,           //优惠券列表
					controller.Coupon.Add,            //添加优惠券
					controller.Coupon.Delete,         //删除优惠券
					controller.Coupon.Update,         //更新优惠券
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
					controller.Goods.ListFrontend,        //商品列表
					controller.Goods.ListByLevelFrontend, //商品列表（根据2级分类）
					controller.Comment.ListFrontend,      //评论列表
					controller.Article.ListFrontend,      //种草推荐文章列表（所有人可见）
				)
				// 需要frontend鉴权的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfUserToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.User.Create,                     //用户注册
						controller.User.Info,                       //用户信息
						controller.User.ResetPassword,              //重置密码
						controller.Order.ListFrontend,              //订单列表（仅用户自己的订单）
						controller.Order.AddFrontend,               //添加订单
						controller.Article.ListMyFrontend,          //文章列表（仅用户自己的文章）
						controller.Article.AddFrontend,             //添加文章
						controller.Article.DeleteFrontend,          //删除文章
						controller.Article.DetailFrontend,          //文章详情
						controller.Praise.ListFrontend,             //点赞列表（仅用户自己的点赞）
						controller.Praise.AddFrontend,              //添加点赞
						controller.Praise.DeleteFrontend,           //删除点赞
						controller.Praise.DeleteByTypeFrontend,     //删除点赞（根据类型）
						controller.Collection.ListFrontend,         //收藏列表（仅用户自己的收藏）
						controller.Collection.AddFrontend,          //添加收藏
						controller.Collection.DeleteFrontend,       //删除收藏
						controller.Collection.DeleteByTypeFrontend, //删除收藏（根据类型）
						controller.Comment.AddFrontend,             //添加评论
						controller.Comment.DeleteFrontend,          //删除评论
						controller.Consignee.ListFrontend,          //查询收货人列表（仅用户自己）
						controller.Consignee.AddFrontend,           //添加收货人
						controller.Consignee.DeleteFrontend,        //删除收货人
						controller.Consignee.UpdateFrontend,        //删除收货人
						controller.Goods.DetailFrontend,            //商品详情
						controller.GoodsOptions.ListFrontend,       //商品规格列表
						controller.GoodsOptions.DetailFrontend,     //商品规格详情
						controller.Cart.ListFrontend,               //购物车列表（仅用户自己）
						controller.Cart.AddFrontend,                //添加购物车
						controller.Cart.DeleteFrontend,             //删除购物车
						controller.Cart.UpdateFrontend,             //更新购物车
						controller.UserCoupon.List,                 //优惠券列表
						controller.UserCoupon.Add,                  //添加优惠券
						controller.UserCoupon.Delete,               //删除优惠券
						controller.UserCoupon.Update,               //更新优惠券
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
