package model

// ConsigneeGetListInput 获取收货人列表
type ConsigneeGetListBackendInput struct {
	Page  int // 分页号码
	Size  int // 分页数量，最大50
	Name  string
	Phone string
}

// ConsigneeGetListInput 获取收货人列表
type ConsigneeGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// ConsigneeGetListOutput 查询列表结果
type ConsigneeGetListOutput struct {
	List  []ConsigneeGetListOutputItem `json:"list" description:"列表"`
	Page  int                          `json:"page" description:"分页码"`
	Size  int                          `json:"size" description:"分页数量"`
	Total int                          `json:"total" description:"数据总数"`
}

type ConsigneeGetListOutputItem struct {
	Id        int    `json:"id"`         // 自增ID
	IsDefault int    `json:"is_default"` // 是否默认
	UserId    int    `json:"user_id"`    // 用户ID
	Name      string `json:"name"`       // 收货人姓名
	Phone     string `json:"phone"`      // 收货人电话
	Province  string `json:"province"`   // 省
	City      string `json:"city"`       // 市
	Town      string `json:"town"`       // 区
	Street    string `json:"street"`     // 街道
	Detail    string `json:"detail"`     // 详细地址
	TimeCommon
}

type ConsigneeCreateUpdateBase struct {
	UserId    int
	IsDefault int
	Name      string
	Phone     string
	Province  string
	City      string
	Town      string
	Street    string
	Detail    string
}

type ConsigneeAddInput struct {
	ConsigneeCreateUpdateBase
}

type ConsigneeAddOutput struct {
	ConsigneeId int
}

type ConsigneeUpdateInput struct {
	ConsigneeCreateUpdateBase
	Id int
}

type ConsigneeUpdateOutput struct{}

type ConsigneeDeleteInput struct {
	Id int `json:"id"`
}

type ConsigneeDeleteOutput struct{}
