package model

type DashboardHeadOutput struct {
	TodayOrderCount int `json:"today_order_count"`
	DAU             int `json:"dau"`
	ConversionRate  int `json:"conversion_rate"`
}
