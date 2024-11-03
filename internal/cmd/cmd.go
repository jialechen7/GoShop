package cmd

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/consts"
	"goshop/internal/controller"
	"goshop/internal/dao"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"goshop/utility"
	"goshop/utility/response"
	"strconv"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/goflyfox/gtoken/gtoken"

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
			gfAdminToken := &gtoken.GfToken{
				ServerName: "goshop",
				//Timeout:         10 * 1000,
				CacheMode:        2,
				LoginPath:        "/backend/login",
				LoginBeforeFunc:  loginBeforeFunc,
				LoginAfterFunc:   loginAfterFunc,
				LogoutPath:       "/backend/logout",
				AuthPaths:        g.SliceStr{"/backend/admin/info", "/backend/admin/update", "/backend/admin/delete"},
				AuthExcludePaths: g.SliceStr{"/backend/admin/add", "/backend/dashboard/head"}, // 不拦截路径
				AuthAfterFunc:    authAfterFunc,
				MultiLogin:       true,
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
					hello.NewV1(),           //示例
					controller.Rotation,     //轮播图
					controller.Position,     //手工位
					controller.Admin.Create, //管理员创建
					controller.Admin.List,   //管理员列表
					controller.Dashboard,    //数据大屏
					controller.Role,         //角色
					controller.Permission,   //权限
					//controller.Login,        //登录（使用gtoken时不需要绑定，在gtoken中绑定）
				)

				// Special handler that needs authentication.
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth)  // 使用gtoken时，不需要绑定jwt中间件
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

// loginBeforeFunc 自定义登录验证
func loginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误."))
		r.ExitAll()
	}
	ctx := context.TODO()

	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("用户名不存在"))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail("密码不正确"))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GtokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// loginAfterFunc 自定义登陆成功后的行为
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump("respData", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		// 此处的userkey为LoginBeforeFunc返回的第二个参数
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtokenAdminPrefix)
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).Where("id", adminId).Scan(&adminInfo)
		if err != nil {
			return
		}

		data := &backend.LoginRes{
			Type:     "Bearer",
			Token:    respData.GetString("token"),
			ExpireAt: 10 * 24 * 60 * 60,
			IsAdmin:  adminInfo.IsAdmin,
			RoleIds:  adminInfo.RoleIds,
		}
		response.JsonExit(r, respData.Code, "", data)
	}
}

// authAfterFunc 自定义验证成功后的行为
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}
