package cmd

import (
	"context"
	"goshop/api/frontend"
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

func StartFrontendGToken() (gfUserToken *gtoken.GfToken, err error) {
	gfUserToken = &gtoken.GfToken{
		ServerName: "goshop-frontend",
		//Timeout:         10 * 1000,
		CacheMode:        consts.CacheModeRedis,
		LoginPath:        "/user/login",
		LoginBeforeFunc:  userLoginBeforeFunc,
		LoginAfterFunc:   userLoginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/frontend/user/*"},
		AuthExcludePaths: g.SliceStr{"/frontend/user/register", "/rotation/*"}, // 不拦截路径
		AuthAfterFunc:    userAuthAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	err = gfUserToken.Start()
	return
}

// loginBeforeFunc 自定义登录验证
func userLoginBeforeFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	userInfo := entity.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, name).Scan(&userInfo)
	if err != nil {
		response.JsonExit(r, consts.UserNameOrPasswordError, consts.ErrUserNotExist, nil)
	}
	if utility.EncryptPassword(password, userInfo.UserSalt) != userInfo.Password {
		response.JsonExit(r, consts.UserNameOrPasswordError, consts.ErrPassword, nil)
	}
	// 唯一标识，扩展参数user data
	return consts.GtokenUserPrefix + strconv.Itoa(userInfo.Id), userInfo
}

// loginAfterFunc 自定义登陆成功后的行为
func userLoginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 1
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 0
		// 此处的userKey为LoginBeforeFunc返回的第二个参数
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GtokenUserPrefix)
		userInfo := entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).Where(dao.UserInfo.Columns().Id, userId).Scan(&userInfo)
		if err != nil {
			return
		}

		data := &frontend.LoginRes{
			Type:     "Bearer",
			Token:    respData.GetString("token"),
			ExpireIn: consts.GtokenExpireIn,
			Name:     userInfo.Name,
			Avatar:   userInfo.Avatar,
			Sign:     userInfo.Sign,
			Sex:      userInfo.Sex,
			Status:   userInfo.Status,
		}
		response.JsonExit(r, respData.Code, "", data)
	}
}

// authAfterFunc 自定义验证成功后的行为
func userAuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.UserInfo
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	if userInfo.Status == consts.UserStatusBlacked {
		response.AuthBlack(r)
		return
	}
	r.SetCtxVar(consts.CtxUserId, userInfo.Id)
	r.SetCtxVar(consts.CtxUserName, userInfo.Name)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserSign, userInfo.Sign)
	r.SetCtxVar(consts.CtxUserSex, userInfo.Sex)
	r.SetCtxVar(consts.CtxUserStatus, userInfo.Status)
	r.Middleware.Next()
}
