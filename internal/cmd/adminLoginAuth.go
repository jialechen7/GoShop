package cmd

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/consts"
	"goshop/internal/dao"
	"goshop/internal/model/entity"
	"goshop/utility"
	"goshop/utility/response"
	"strconv"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		ServerName: "goshop-backend",
		//Timeout:         10 * 1000,
		CacheMode:        consts.CacheModeRedis,
		LoginPath:        "/admin/login",
		LoginBeforeFunc:  loginBeforeFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/admin/logout",
		AuthPaths:        g.SliceStr{"/"},
		AuthExcludePaths: g.SliceStr{"/backend/admin/add"}, // 不拦截路径
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	err = gfAdminToken.Start()
	return
}

// loginBeforeFunc 自定义登录验证
func loginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, name).Scan(&adminInfo)
	if err != nil {
		response.JsonExit(r, consts.UserNameOrPasswordError, consts.ErrUserNotExist, nil)
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		response.JsonExit(r, consts.UserNameOrPasswordError, consts.ErrPassword, nil)
	}
	// 唯一标识，扩展参数user data
	return consts.GtokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// loginAfterFunc 自定义登陆成功后的行为
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 1
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 0
		// 此处的userKey为LoginBeforeFunc返回的第二个参数
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtokenAdminPrefix)
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).Where(dao.AdminInfo.Columns().Id, adminId).Scan(&adminInfo)
		if err != nil {
			return
		}

		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}

		permissionsIds := g.Slice{}
		for _, rolePermissionInfo := range rolePermissionInfos {
			permissionsIds = append(permissionsIds, rolePermissionInfo.PermissionId)
		}

		permissions := make([]entity.PermissionInfo, 0)
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionsIds).Scan(&permissions)
		if err != nil {
			return
		}

		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GtokenExpireIn,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
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
