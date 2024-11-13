package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ConsigneeGetListCommonReq struct {
	g.Meta `path:"/consignee/list" tags:"Consignee" method:"get" summary:"收货人列表接口"`
	CommonPaginationReq
}

type ConsigneeGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type ConsigneeAddReq struct {
	g.Meta    `path:"/consignee/add" tags:"Consignee" method:"post" summary:"创建收货人接口"`
	Name      string `json:"name" form:"name" v:"required#请输入收货人姓名" dc:"收货人姓名"`
	IsDefault int    `json:"is_default" form:"is_default" v:"required#请选择是否默认" dc:"是否默认"`
	Phone     string `json:"phone" form:"phone" v:"required#请输入手机号" dc:"手机号"`
	Province  string `json:"province" form:"province" v:"required#请选择省份" dc:"省份"`
	City      string `json:"city" form:"city" v:"required#请选择城市" dc:"城市"`
	Town      string `json:"town" form:"town" v:"required#请选择区县" dc:"区县"`
	Street    string `json:"street" form:"street" v:"required#请输入街道" dc:"街道"`
	Detail    string `json:"detail" form:"detail" v:"required#请输入详细地址" dc:"详细地址"`
}

type ConsigneeAddRes struct {
	ConsigneeId int `json:"consignee_id"`
}

type ConsigneeDeleteReq struct {
	g.Meta `path:"/consignee/delete" method:"delete" tags:"Consignee" summary:"删除收货人接口"`
	Id     int `v:"min:1#请选择需要删除的收货人" dc:"收货人id"`
}
type ConsigneeDeleteRes struct{}

type ConsigneeUpdateReq struct {
	g.Meta    `path:"/consignee/update" method:"post" tags:"收货人" summary:"修改收货人接口"`
	Id        int    `json:"id" form:"id" v:"required#请输入id" dc:"收货人id"`
	Name      string `json:"name" form:"name" v:"required#请输入收货人姓名" dc:"收货人姓名"`
	IsDefault int    `json:"is_default" form:"is_default" v:"required#请选择是否默认" dc:"是否默认"`
	Phone     string `json:"phone" form:"phone" v:"required#请输入手机号" dc:"手机号"`
	Province  string `json:"province" form:"province" v:"required#请选择省份" dc:"省份"`
	City      string `json:"city" form:"city" v:"required#请选择城市" dc:"城市"`
	Town      string `json:"town" form:"town" v:"required#请选择区县" dc:"区县"`
	Street    string `json:"street" form:"street" v:"required#请输入街道" dc:"街道"`
	Detail    string `json:"detail" form:"detail" v:"required#请输入详细地址" dc:"详细地址"`
}
type ConsigneeUpdateRes struct{}
