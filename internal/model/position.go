package model

// PositionCreateUpdateBase 创建/修改手工位基类
type PositionCreateUpdateBase struct {
	PicUrl    string
	Link      string
	GoodsName string
	GoodsId   int
	Sort      int
}

// PositionCreateInput 创建手工位
type PositionCreateInput struct {
	PositionCreateUpdateBase
}

// PositionCreateOutput 创建手工位返回结果
type PositionCreateOutput struct {
	PositionId int `json:"position_id"`
}

// PositionUpdateInput 修改手工位
type PositionUpdateInput struct {
	PositionCreateUpdateBase
	Id int
}

// PositionGetListInput 获取手工位列表
type PositionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型
}

// PositionGetListOutput 查询列表结果
type PositionGetListOutput struct {
	List  []PositionGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

type PositionGetListOutputItem struct {
	//Position  *PositionListItem `json:"Position"`
	Id        int    `json:"id"` // 自增ID
	PicUrl    string `json:"pic_url"`
	Link      string `json:"link"`
	GoodsName string `json:"goods_name"`
	GoodsId   int    `json:"goods_id"`
	Sort      int    `json:"sort"` // 排序，数值越高越靠前
	TimeCommon
}

//// PositionListItem 主要用于手工位列表展示
//type PositionListItem struct {
//	Id        int         `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_url"`
//	Link      string      `json:"link"`
//	GoodsName string      `json:"goods_name"`
//	GoodsId   int         `json:"goods_id"`
//	Sort      int         `json:"sort"`       // 排序，数值越高越靠前
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}
