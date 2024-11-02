package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/service"
)

type sDashboard struct{}

var Dashboard = sDashboard{}

func (s *sDashboard) DashboardHead(ctx context.Context, req *backend.DashboardHeadReq) (out *backend.DashboardHeadRes, err error) {
	res, err := service.Dashboard().DashboardHead(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DashboardHeadRes{
		TodayOrderCount: res.TodayOrderCount,
		DAU:             res.DAU,
		ConversionRate:  res.ConversionRate,
	}, nil
}
