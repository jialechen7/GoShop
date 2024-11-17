package model

// GoodsGetListInput 获取商品列表
type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// GoodsGetListByLevelInput 获取商品列表
type GoodsGetListByLevelInput struct {
	Page    int // 分页号码
	Size    int // 分页数量，最大50
	LevelId int // 分类ID
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

type GoodsGetListOutputItem struct {
	Id               int    `json:"id"`                 // 自增ID
	PicUrl           string `json:"pic_url"`            // 图片地址
	Name             string `json:"name"`               // 名称
	Price            int    `json:"price"`              // 价格
	Level1CategoryId int    `json:"level1_category_id"` // 一级分类ID
	Level2CategoryId int    `json:"level2_category_id"` // 二级分类ID
	Level3CategoryId int    `json:"level3_category_id"` // 三级分类ID
	Brand            string `json:"brand"`              // 品牌
	Stock            int    `json:"stock"`              // 库存
	Sale             int    `json:"sale"`               // 销量
	Tags             string `json:"tags"`               // 标签
	DetailInfo       string `json:"detail_info"`        // 详情
	TimeCommon
}

type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

type GoodsAddInput struct {
	GoodsCreateUpdateBase
}

type GoodsAddOutput struct {
	GoodsId int
}

type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id int
}

type GoodsUpdateOutput struct{}

type GoodsDeleteInput struct {
	Id int `json:"id"`
}

type GoodsDeleteOutput struct{}

type GoodsDetailInput struct {
	Id int
}

type GoodsDetailOutput struct {
	Id               int    `json:"id"`
	PicUrl           string `json:"pic_url"`
	Name             string `json:"name"`
	Price            int    `json:"price"`
	Level1CategoryId int    `json:"level1_category_id"`
	Level2CategoryId int    `json:"level2_category_id"`
	Level3CategoryId int    `json:"level3_category_id"`
	Brand            string `json:"brand"`
	Stock            int    `json:"stock"`
	Sale             int    `json:"sale"`
	Tags             string `json:"tags"`
	DetailInfo       string `json:"detail_info"`
	IsPraise         int    `json:"is_praise"`
	IsCollect        int    `json:"is_collect"`
	TimeCommon
}
