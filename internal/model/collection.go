package model

// CollectionGetListInput 获取收藏列表
type CollectionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Type int // 类型
}

// CollectionGetListOutput 查询列表结果
type CollectionGetListOutput struct {
	List  []CollectionGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

type CollectionGetListOutputItem struct {
	Id          int          `json:"id"` // 自增ID
	UserId      int          `json:"user_id"`
	Type        int          `json:"type"`
	ObjectId    int          `json:"object_id"`
	ArticleInfo *ArticleInfo `json:"article_info" orm:"with:id=object_id"`
	GoodsInfo   *GoodsInfo   `json:"goods_info" orm:"with:id=object_id"`
	TimeCommon
}

type CollectionCreateUpdateBase struct {
	UserId   int
	Type     int
	ObjectId int
}

type CollectionAddInput struct {
	CollectionCreateUpdateBase
}

type CollectionAddOutput struct {
	CollectionId int
}

type CollectionDeleteInput struct {
	Id int `json:"id"`
}

type CollectionDeleteOutput struct{}

type CollectionDeleteByTypeInput struct {
	Type     int `json:"type"`
	UserId   int `json:"user_id"`
	ObjectId int `json:"object_id"`
}

type CollectionDeleteByTypeOutput struct{}
