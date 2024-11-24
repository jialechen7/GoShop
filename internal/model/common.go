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
	Id         int         `json:"id"` // 自增ID
	Title      string      `json:"title"`
	Desc       string      `json:"desc"`
	Detail     string      `json:"detail"`
	PicUrl     string      `json:"pic_url"`
	IsAdmin    int         `json:"is_admin"`
	CreatedAt  *gtime.Time `json:"created_at"`
	UpdatedAt  *gtime.Time `json:"updated_at"`
}

type GoodsInfo struct {
	gmeta.Meta       `orm:"table:goods_info"`
	Id               int         `json:"id"`
	PicUrl           string      `json:"pic_url"`
	Name             string      `json:"name"`
	Price            float64     `json:"price"`
	Level1CategoryId int         `json:"level1_category_id"`
	Level2CategoryId int         `json:"level2_category_id"`
	Level3CategoryId int         `json:"level3_category_id"`
	Brand            string      `json:"brand"`
	Stock            int         `json:"stock"`
	Sale             int         `json:"sale"`
	Tags             string      `json:"tags"`
	DetailInfo       string      `json:"detail_info"`
	CreatedAt        *gtime.Time `json:"created_at"`
	UpdatedAt        *gtime.Time `json:"updated_at"`
}

type GoodsOptionsInfo struct {
	gmeta.Meta `orm:"table:goods_options_info"`
	Id         int         `json:"id"`
	GoodsId    int         `json:"goods_id"`
	PicUrl     string      `json:"pic_url"`
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	Stock      int         `json:"stock"`
	GoodsInfo  *GoodsInfo  `json:"goods_info" orm:"with:id=goods_id"`
	CreatedAt  *gtime.Time `json:"created_at"`
	UpdatedAt  *gtime.Time `json:"updated_at"`
}

type CouponInfo struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Condition  int         `json:"condition"`
	Price      int         `json:"price"`
	GoodsIds   string      `json:"goods_ids"`
	CategoryId int         `json:"category_id"`
	Type       int         `json:"type"`
	CreatedAt  *gtime.Time `json:"created_at"`
	UpdatedAt  *gtime.Time `json:"updated_at"`
}

type SeckillCouponInfo struct {
	gmeta.Meta `orm:"table:seckill_coupon_info"`
	Id         int         `json:"id"`
	CouponId   int         `json:"coupon_id"`
	Stock      int         `json:"stock"`
	StartTime  *gtime.Time `json:"start_time"`
	EndTime    *gtime.Time `json:"end_time"`
	CreatedAt  *gtime.Time `json:"created_at"`
	UpdatedAt  *gtime.Time `json:"updated_at"`
}
