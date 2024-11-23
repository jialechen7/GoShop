package oauth

import (
	"context"
	"database/sql"
	"errors"
	"goshop/internal/cmd"
	"goshop/internal/consts"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"goshop/utility"
	"net/url"
	"strconv"
	"strings"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/go-sql-driver/mysql"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sOauth struct{}

func init() {
	service.RegisterOauth(New())
}

func New() *sOauth {
	return &sOauth{}
}

func setCache(gfToken *gtoken.GfToken, ctx context.Context, cacheKey string, userCache g.Map) (resp gtoken.Resp) {
	switch gfToken.CacheMode {
	case consts.CacheModeRedis:
		cacheValueJson, err1 := gjson.Encode(userCache)
		if err1 != nil {
			g.Log().Error(ctx, "[GToken]cache json encode error", err1)
			return gtoken.Error("cache json encode error")
		}
		_, err := g.Redis().Do(ctx, "SETEX", cacheKey, gfToken.Timeout/1000, cacheValueJson)
		if err != nil {
			g.Log().Error(ctx, "[GToken]cache set error", err)
			return gtoken.Error("cache set error")
		}
	default:
		return gtoken.Error("cache model error")
	}

	return gtoken.Succ(userCache)
}

func (s *sOauth) GithubLogin(ctx context.Context) (string, error) {
	clientId := g.Cfg().MustGet(ctx, "oauth.github.client_id").String()
	redirectUri := g.Cfg().MustGet(ctx, "oauth.github.redirect_uri").String()
	state := grand.S(10)
	_, err := g.Redis().Do(ctx, "SET", consts.GithubStatePrefix+state, state, "EX", consts.StateExpireIn)
	if err != nil {
		return "", err
	}
	url := g.Cfg().MustGet(ctx, "oauth.github.authorize_url").String() + "?client_id=" + clientId + "&redirect_uri=" + url.QueryEscape(redirectUri) + "&state=" + state
	return url, nil
}

// GithubReceiveCode 接收github回调并自动登录（未注册自动注册）
func (s *sOauth) GithubReceiveCode(ctx context.Context, in model.GithubReceiveCodeInput) (out model.GithubReceiveCodeOutput, err error) {
	// 1. 查询回调带有的state是否与发送的state一致，不一致则返回错误
	gvar, err := g.Redis().Do(ctx, "GET", consts.GithubStatePrefix+in.State)
	if err != nil {
		return model.GithubReceiveCodeOutput{}, err
	}
	if gvar == nil || gconv.String(gvar) != in.State {
		return model.GithubReceiveCodeOutput{}, gerror.New(consts.ErrState)
	}
	// 2. 配置参数请求github的access_token
	clientId := g.Cfg().MustGet(ctx, "oauth.github.client_id").String()
	clientSecret := g.Cfg().MustGet(ctx, "oauth.github.client_secret").String()
	redirectUri := g.Cfg().MustGet(ctx, "oauth.github.redirect_uri").String()
	code := in.Code

	client := g.Client().Proxy(g.Cfg().MustGet(ctx, "proxy").String())
	tokenResponse := client.PostContent(ctx, g.Cfg().MustGet(ctx, "oauth.github.access_token_url").String(), g.Map{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
		"redirect_uri":  redirectUri,
	})

	// 3. 解析返回的access_token
	parts := strings.Split(tokenResponse, "&")
	params := make(map[string]string)

	for _, part := range parts {
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			params[keyValue[0]] = keyValue[1]
		}
	}
	accssToken := params["access_token"]

	// 4. 请求用户信息
	headerMap := make(map[string]string)
	headerMap["Authorization"] = "Bearer " + accssToken
	userInfo := client.Header(headerMap).GetVar(ctx, g.Cfg().MustGet(ctx, "oauth.github.user_info_url").String())
	userMap := userInfo.Map()
	openId := gconv.String(userMap["id"])
	name := userMap["login"].(string)

	// 5. 判断是否已经注册
	adminEntity := &entity.AdminInfo{}
	err = dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().GithubOpenid, openId).Scan(&adminEntity)
	if !errors.Is(err, sql.ErrNoRows) && err != nil {
		return model.GithubReceiveCodeOutput{}, err
	}

	// 6. 用户不存在，先创建
	if errors.Is(err, sql.ErrNoRows) {
		adminEntity.Name = name
		adminEntity.GithubOpenid = openId
		adminEntity.UserSalt = gconv.String(grand.S(10))
		adminEntity.Password = utility.EncryptPassword(consts.AdminDefaultPwd, adminEntity.UserSalt)
		id, err := dao.AdminInfo.Ctx(ctx).Data(adminEntity).OmitEmpty().InsertAndGetId()
		var mysqlErr *mysql.MySQLError
		// Duplicate entry 错误码为1062
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			adminEntity.Name += "#" + grand.S(4)
			id, err = dao.AdminInfo.Ctx(ctx).Data(adminEntity).OmitEmpty().InsertAndGetId()
			if err != nil {
				return model.GithubReceiveCodeOutput{}, err
			}
		} else if err != nil {
			return model.GithubReceiveCodeOutput{}, err
		}
		adminEntity.Id = int(id)
	}

	userKey := consts.GtokenAdminPrefix + strconv.Itoa(adminEntity.Id)

	// 这块就是gtoken的genToken方法，但是由于私有化，需要自己写一个
	token := cmd.GfAdminToken.EncryptToken(ctx, userKey, "")
	if !token.Success() {
		return model.GithubReceiveCodeOutput{}, errors.New(consts.ErrEncryptToken)
	}

	cacheKey := cmd.GfAdminToken.CacheKey + userKey
	userCache := g.Map{
		gtoken.KeyUserKey:     userKey,
		gtoken.KeyUuid:        token.GetString(gtoken.KeyUuid),
		gtoken.KeyData:        adminEntity,
		gtoken.KeyCreateTime:  gtime.Now().TimestampMilli(),
		gtoken.KeyRefreshTime: gtime.Now().TimestampMilli() + gconv.Int64(cmd.GfAdminToken.MaxRefresh),
	}
	cacheResp := setCache(cmd.GfAdminToken, ctx, cacheKey, userCache)
	if !cacheResp.Success() {
		return model.GithubReceiveCodeOutput{}, errors.New(cacheResp.Msg)
	}
	// 这块就是gtoken的genToken方法，但是由于私有化，需要自己写一个

	out.Type = "Bearer"
	out.Token = token.GetString("token")
	out.ExpireIn = consts.GtokenExpireIn
	out.IsAdmin = adminEntity.IsAdmin
	out.RoleIds = adminEntity.RoleIds
	// 7. 根据用户的角色获取权限
	var rolePermissionInfos []entity.RolePermissionInfo
	err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminEntity.RoleIds}).Scan(&rolePermissionInfos)
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

	out.Permissions = permissions
	return out, nil
}
