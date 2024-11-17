package comment

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

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

// GetListBackend 查询评论列表
func (s *sComment) GetListBackend(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error) {
	var (
		m = dao.CommentInfo.Ctx(ctx)
	)
	out = &model.CommentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CommentInfo
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

// GetListFrontend 查询评论列表
func (s *sComment) GetListFrontend(ctx context.Context, in model.CommentGetListFrontendInput) (out *model.CommentGetListOutput, err error) {
	var (
		m = dao.CommentInfo.Ctx(ctx)
	)
	out = &model.CommentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.Where(g.Map{
		dao.CommentInfo.Columns().Type:     in.Type,
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
	})
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.CommentInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// UpdateBackend 更新评论
func (s *sComment) UpdateBackend(ctx context.Context, in model.CommentUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return err
	}
	_, err := dao.CommentInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.CommentInfo.Columns().Id, in.Id).Update()
	return err
}

// DeleteBackend 删除评论
func (s *sComment) DeleteBackend(ctx context.Context, id int) error {
	return dao.CommentInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除评论
		_, err := dao.CommentInfo.Ctx(ctx).Where(g.Map{
			dao.CommentInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// AddFrontend 添加评论
func (s *sComment) AddFrontend(ctx context.Context, in model.CommentAddInput) (out *model.CommentAddOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.CommentInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.CommentAddOutput{CommentId: int(lastInsertID)}, err
}

// DeleteFrontend 删除评论
func (s *sComment) DeleteFrontend(ctx context.Context, id int) error {
	var commentInfo entity.CommentInfo
	err := dao.CommentInfo.Ctx(ctx).Where(dao.CommentInfo.Columns().Id, id).Scan(&commentInfo)
	if err != nil {
		return err
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	if userId != commentInfo.UserId {
		return gerror.New(consts.ErrNoPermission)
	}

	return dao.CommentInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除评论
		_, err := dao.CommentInfo.Ctx(ctx).Where(g.Map{
			dao.CommentInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}
