package dashboard

import (
	"context"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sDashboard struct{}

func init() {
	service.RegisterDashboard(New())
}

func New() *sDashboard {
	return &sDashboard{}
}

func (s *sDashboard) DashboardHead(ctx context.Context) (out *model.DashboardHeadOutput, err error) {
	return &model.DashboardHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             dau(ctx),
		ConversionRate:  conversionRate(ctx),
	}, nil
}

func todayOrderCount(ctx context.Context) int {
	startOfDay := gtime.Now().StartOfDay()
	endOfDay := gtime.Now().EndOfDay()
	orderCount, err := dao.OrderInfo.Ctx(ctx).WhereBetween(dao.OrderInfo.Columns().CreatedAt, startOfDay, endOfDay).Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return orderCount
}

func dau(ctx context.Context) int {
	return grand.Intn(1000)
}

func conversionRate(ctx context.Context) int {
	return grand.Intn(100)
}
