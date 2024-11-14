package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type TimeCommon struct {
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type ArticleInfo struct {
	gmeta.Meta `orm:"table:article_info"`
	Id         int    `json:"id"` // 自增ID
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Detail     string `json:"detail"`
	PicUrl     string `json:"pic_url"`
	IsAdmin    int    `json:"is_admin"`
	TimeCommon
}

type GoodsInfo struct {
	gmeta.Meta       `orm:"table:goods_info"`
	Id               int     `json:"id"`
	PicUrl           string  `json:"pic_url"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	Level1CategoryId int     `json:"level1_category_id"`
	Level2CategoryId int     `json:"level2_category_id"`
	Level3CategoryId int     `json:"level3_category_id"`
	Brand            string  `json:"brand"`
	Stock            int     `json:"stock"`
	Sale             int     `json:"sale"`
	Tags             string  `json:"tags"`
	DetailInfo       string  `json:"detail_info"`
	TimeCommon
}

type GoodsOptionsInfo struct {
	gmeta.Meta `orm:"table:goods_options_info"`
	Id         int        `json:"id"`
	GoodsId    int        `json:"goods_id"`
	PicUrl     string     `json:"pic_url"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Stock      int        `json:"stock"`
	GoodsInfo  *GoodsInfo `json:"goods_info" orm:"with:id=goods_id"`
	TimeCommon
}
