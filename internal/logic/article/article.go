package article

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
)

type sArticle struct{}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

// GetListBackend 查询文章列表
func (s *sArticle) GetListBackend(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error) {
	var (
		m = dao.ArticleInfo.Ctx(ctx)
	)
	out = &model.ArticleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ArticleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// UpdateBackend 更新文章
func (s *sArticle) UpdateBackend(ctx context.Context, in model.ArticleUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	_, err := dao.ArticleInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.ArticleInfo.Columns().Id, in.Id).Update()
	return err
}

// AddBackend 添加文章
func (s *sArticle) AddBackend(ctx context.Context, in model.ArticleAddInput) (out *model.ArticleAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.ArticleInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.ArticleAddOutput{ArticleId: int(lastInsertID)}, err
}

// DeleteBackend 删除文章
func (s *sArticle) DeleteBackend(ctx context.Context, id int) error {
	return dao.ArticleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除文章
		_, err := dao.ArticleInfo.Ctx(ctx).Where(g.Map{
			dao.ArticleInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// GetListFrontend 查询文章列表（仅用户发表的文章）
func (s *sArticle) GetListFrontend(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error) {
	var (
		m = dao.ArticleInfo.Ctx(ctx)
	)
	out = &model.ArticleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Where(dao.ArticleInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxUserId))).Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.ArticleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// AddFrontend 添加文章
func (s *sArticle) AddFrontend(ctx context.Context, in model.ArticleAddInput) (out *model.ArticleAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.ArticleInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.ArticleAddOutput{ArticleId: int(lastInsertID)}, err
}

// DeleteFrontend 删除文章
func (s *sArticle) DeleteFrontend(ctx context.Context, id int) error {
	var articleInfo entity.ArticleInfo
	err := dao.ArticleInfo.Ctx(ctx).Where(dao.ArticleInfo.Columns().Id, id).Scan(&articleInfo)
	if err != nil {
		return err
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	// 判断当前用户是否有权限对该文章操作
	if userId != articleInfo.UserId {
		return gerror.New(consts.ErrNoPermission)
	}

	return dao.ArticleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除文章
		_, err := dao.ArticleInfo.Ctx(ctx).Where(g.Map{
			dao.ArticleInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// DetailFrontend 查询文章详情
func (s *sArticle) DetailFrontend(ctx context.Context, id int) (out *model.ArticleDetailOutput, err error) {
	var articleInfo entity.ArticleInfo
	err = dao.ArticleInfo.Ctx(ctx).Where(dao.ArticleInfo.Columns().Id, id).Scan(&articleInfo)
	if err != nil {
		return nil, err
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	articleId := articleInfo.Id
	queryType := consts.PraiseArticleType

	// 判断当前用户是否对该文章点赞
	isPraise, _ := dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		dao.PraiseInfo.Columns().Type:     queryType,
		dao.PraiseInfo.Columns().UserId:   userId,
		dao.PraiseInfo.Columns().ObjectId: articleId,
	}).Count()

	return &model.ArticleDetailOutput{
		Id:       articleId,
		UserId:   articleInfo.UserId,
		Title:    articleInfo.Title,
		Desc:     articleInfo.Desc,
		Detail:   articleInfo.Detail,
		PicUrl:   articleInfo.PicUrl,
		Praise:   articleInfo.Praise,
		IsPraise: isPraise,
		IsAdmin:  articleInfo.IsAdmin,
		TimeCommon: model.TimeCommon{
			CreatedAt: articleInfo.CreatedAt.String(),
			UpdatedAt: articleInfo.UpdatedAt.String(),
		},
	}, nil
}
