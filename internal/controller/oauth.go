package controller

import (
	"context"
	"fmt"
	"goshop/api/backend"
	"goshop/internal/model"
	"goshop/internal/service"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

var Oauth = cOauth{}

type cOauth struct{}

func (c *cOauth) GithubLogin(ctx context.Context, req *backend.GithubLoginReq) (res *backend.OauthLoginRes, err error) {
	url, err := service.Oauth().GithubLogin(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.OauthLoginRes{Url: url}, nil
}

func (c *cOauth) GithubReceiveCode(r *ghttp.Request) {
	req := &backend.GithubReceiveCodeReq{}
	if err := r.Parse(req); err != nil {
		r.Response.RedirectTo("http://127.0.0.1:8080/#/login")
	}
	ctx := context.Background()
	out, err := service.Oauth().GithubReceiveCode(ctx, model.GithubReceiveCodeInput{
		Code:  req.Code,
		State: req.State,
	})
	if err != nil {
		r.Response.RedirectTo("http://127.0.0.1:8080/#/login")
	}

	g.Dump(out.Token)
	r.Response.RedirectTo(fmt.Sprintf("http://127.0.0.1:8080/#/login?type=%s&token=%s&expire_in=%d&is_admin=%d&role_ids=%s", out.Type, url.QueryEscape(out.Token), out.ExpireIn, out.IsAdmin, out.RoleIds))
}
