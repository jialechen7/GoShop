package coupon

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/util/gconv"

	"goshop/internal/dao"
	"goshop/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sSeckillCoupon struct{}

func init() {
	service.RegisterSeckillCoupon(New())
}

func New() *sSeckillCoupon {
	return &sSeckillCoupon{}
}

// GetList 查询秒杀优惠券列表
func (s *sSeckillCoupon) GetList(ctx context.Context, in model.SeckillCouponGetListInput) (out *model.SeckillCouponGetListOutput, err error) {
	var (
		m = dao.SeckillCouponInfo.Ctx(ctx)
	)
	out = &model.SeckillCouponGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m = m.OrderAsc(dao.SeckillCouponInfo.Columns().StartTime)
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.SeckillCouponInfo
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

// Add 添加秒杀优惠券
func (s *sSeckillCoupon) Add(ctx context.Context, in model.SeckillCouponAddInput) (out *model.SeckillCouponAddOutput, err error) {
	couponAddInput := model.CouponAddInput{}
	couponAddInput.Name = in.Name
	couponAddInput.Condition = in.Condition
	couponAddInput.Price = in.Price
	couponAddInput.GoodsIds = in.GoodsIds
	couponAddInput.CategoryId = in.CategoryId
	couponAddInput.Type = consts.CouponTypeSeckill
	couponId, err := dao.CouponInfo.Ctx(ctx).OmitEmpty().Data(couponAddInput).InsertAndGetId()
	if err != nil {
		return out, err
	}
	in.SeckillCouponCreateUpdateBase.CouponId = int(couponId)
	lastInsertID, err := dao.SeckillCouponInfo.Ctx(ctx).OmitEmpty().Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	_, err = g.Redis().Do(ctx, "set", consts.SeckillStockRedisPrefix+gconv.String(couponId), in.Stock)
	if err != nil {
		return nil, err
	}
	return &model.SeckillCouponAddOutput{SeckillCouponId: int(lastInsertID)}, err
}

// Delete 删除秒杀优惠券
func (s *sSeckillCoupon) Delete(ctx context.Context, id int) error {
	return dao.SeckillCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除秒杀优惠券
		_, err := dao.SeckillCouponInfo.Ctx(ctx).Where(g.Map{
			dao.SeckillCouponInfo.Columns().Id: id,
		}).Unscoped().Delete()
		return err
	})
}

// Update 更新秒杀优惠券
func (s *sSeckillCoupon) Update(ctx context.Context, in model.SeckillCouponUpdateInput) error {
	_, err := dao.SeckillCouponInfo.Ctx(ctx).Where(dao.SeckillCouponInfo.Columns().Id, in.Id).OmitEmpty().Data(in).Update()
	return err
}

// Kill 用户秒杀优惠券
func (s *sSeckillCoupon) Kill(ctx context.Context, couponId int) error {
	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	gRes, err := g.Redis().Eval(ctx, consts.LuaSeckillScript, 3,
		[]string{consts.SeckillStockRedisPrefix, consts.UserHasSeckillRedisPrefix, consts.StreamUserCoupon}, []interface{}{couponId, userId})
	if err != nil {
		return err
	}
	if v := gRes.Int(); v != 0 {
		if v == 1 {
			return gerror.New(consts.ErrStockNotEnough)
		} else if v == 2 {
			return gerror.New(consts.ErrHasSeckill)
		}
		return gerror.New(consts.ErrSeckill)
	}

	return nil
}

type StreamData struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

func parseStreamData(rawData []interface{}) *StreamData {
	if len(rawData) == 0 {
		return nil
	}
	streamData := &StreamData{}
	streamData.ID = gconv.String(rawData[0])
	data := rawData[1].([]interface{})
	streamData.Data = make(map[string]interface{})
	for i := 0; i < len(data); i += 2 {
		streamData.Data[gconv.String(data[i])] = data[i+1]
	}
	return streamData
}

// 定义一个处理流消息的公共函数
func processCouponMessage(ctx context.Context, streamReadMode string) error {
	// 1. 获取消息队列中的信息
	gMessage, err := g.Redis().Do(ctx, "XREADGROUP", "GROUP",
		consts.StreamUserCouponGroup, consts.StreamUserCouponConsumer,
		"COUNT", consts.StreamUserCouponOnceRead, "BLOCK", consts.StreamUserCouponBlock,
		"STREAMS", consts.StreamUserCoupon, streamReadMode)
	if err != nil {
		return err
	}

	// 2. 判断是否获取成功
	messageList := gMessage.Slice()
	if len(messageList) == 0 {
		return nil
	}

	// 3. 解析数据
	record := messageList[0].(map[interface{}]interface{})
	rawData := record[consts.StreamUserCoupon].([]interface{})[0].([]interface{})
	streamData := parseStreamData(rawData)

	// 4. 创建用户优惠券（扣减库存，创建用户优惠券）
	err = dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.SeckillCouponInfo.Ctx(ctx).
			Where(dao.SeckillCouponInfo.Columns().CouponId, gconv.Int(streamData.Data["coupon_id"])).
			Decrement(dao.SeckillCouponInfo.Columns().Stock, 1)
		if err != nil {
			return err
		}

		_, err = dao.UserCouponInfo.Ctx(ctx).Data(g.Map{
			dao.UserCouponInfo.Columns().UserId:   gconv.Int(streamData.Data["user_id"]),
			dao.UserCouponInfo.Columns().CouponId: gconv.Int(streamData.Data["coupon_id"]),
			dao.UserCouponInfo.Columns().Status:   consts.CouponStatusAvailable,
		}).Insert()
		return err
	})
	if err != nil {
		return err
	}

	// 5. 确认消息
	_, err = g.Redis().Do(ctx, "XACK", consts.StreamUserCoupon, consts.StreamUserCouponGroup, streamData.ID)
	return err
}

// UserCouponStreamConsumer 用户优惠券消息队列消费者
func UserCouponStreamConsumer(ctx context.Context) {
	for {
		// 使用公共函数处理流消息
		err := processCouponMessage(ctx, consts.StreamReadLatest)
		if err != nil {
			UserCouponPendingListHandler(ctx) // 如果处理失败，转到 Pending List
			continue
		}
	}
}

// UserCouponPendingListHandler 用户优惠券消息队列待处理消息消费者
func UserCouponPendingListHandler(ctx context.Context) {
	for {
		// 使用公共函数处理 Pending List 中的消息
		err := processCouponMessage(ctx, consts.StreamReadPendingList)
		if err != nil {
			// 如果获取消息失败或者处理失败，跳出循环
			return
		}
	}
}
