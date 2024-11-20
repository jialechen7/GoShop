package backend

import "github.com/gogf/gf/v2/frame/g"

type DashboardHeadReq struct {
	g.Meta `path:"/dashboard/head" tags:"数据大屏后台" method:"get" summary:"数据大屏头部信息接口"`
}

type DashboardHeadRes struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单数"`
	DAU             int `json:"dau" dc:"日活用户"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}
