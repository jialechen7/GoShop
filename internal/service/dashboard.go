// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model"
)

type (
	IDashboard interface {
		DashboardHead(ctx context.Context) (out *model.DashboardHeadOutput, err error)
	}
)

var (
	localDashboard IDashboard
)

func Dashboard() IDashboard {
	if localDashboard == nil {
		panic("implement not found for interface IDashboard, forgot register?")
	}
	return localDashboard
}

func RegisterDashboard(i IDashboard) {
	localDashboard = i
}
