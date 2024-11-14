package model

// GoodsOptionsGetListInput 获取商品规格列表
type GoodsOptionsGetListInput struct {
	Page    int // 分页号码
	Size    int // 分页数量，最大50
	GoodsId int // 商品id
}

// GoodsOptionsGetListOutput 查询列表结果
type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsGetListOutputItem `json:"list" description:"列表"`
	Page  int                             `json:"page" description:"分页码"`
	Size  int                             `json:"size" description:"分页数量"`
	Total int                             `json:"total" description:"数据总数"`
}

type GoodsOptionsGetListOutputItem struct {
	Id      int    `json:"id"`       // 自增ID
	GoodsId int    `json:"goods_id"` // 商品ID
	PicUrl  string `json:"pic_url"`  // 图片地址
	Name    string `json:"name"`     // 名称
	Price   int    `json:"price"`    // 价格
	Stock   int    `json:"stock"`    // 库存
	TimeCommon
}

type GoodsOptionsCreateUpdateBase struct {
	GoodsId int
	PicUrl  string
	Name    string
	Price   int
	Stock   int
}

type GoodsOptionsAddInput struct {
	GoodsOptionsCreateUpdateBase
}

type GoodsOptionsAddOutput struct {
	GoodsOptionsId int
}

type GoodsOptionsUpdateInput struct {
	GoodsOptionsCreateUpdateBase
	Id int
}

type GoodsOptionsUpdateOutput struct{}

type GoodsOptionsDeleteInput struct {
	Id int `json:"id"`
}

type GoodsOptionsDeleteOutput struct{}

type GoodsOptionsDetailInput struct {
	Id int
}

type GoodsOptionsDetailOutput struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goods_id"`
	PicUrl  string `json:"pic_url"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Stock   int    `json:"stock"`
	TimeCommon
}
