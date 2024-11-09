package model

// PermissionCreateUpdateBase 创建/修改权限基类
type PermissionCreateUpdateBase struct {
	Name string
	Path string
}

// PermissionCreateInput 创建权限
type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

// PermissionCreateOutput 创建权限返回结果
type PermissionCreateOutput struct {
	PermissionId int `json:"permission_id"`
}

// PermissionUpdateInput 修改权限
type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id int
}

// PermissionGetListInput 获取权限列表
type PermissionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// PermissionGetListOutput 查询列表结果
type PermissionGetListOutput struct {
	List  []PermissionGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

type PermissionGetListOutputItem struct {
	//Permission  *PermissionListItem `json:"Permission"`
	Id   int    `json:"id"`   // 自增ID
	Name string `json:"name"` // 权限名
	Path string `json:"path"`
	TimeCommon
}
