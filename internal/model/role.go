package model

type RoleCreateUpdateBase struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type RoleCreateInput struct {
	RoleCreateUpdateBase
}

type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id int `json:"id"`
}

type RoleUpdateOutput struct {
	RoleId int `json:"role_id"`
}

// RoleGetListInput 获取角色列表
type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

// RoleGetListOutput 查询列表结果
type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

type RoleGetListOutputItem struct {
	//Role  *RoleListItem `json:"Role"`
	Id   int    `json:"id"` // 自增ID
	Name string `json:"name"`
	Desc string `json:"desc"`
	TimeCommon
}

type RoleAddPermissionsInput struct {
	RoleId        int   `json:"role_id"`
	PermissionIds []int `json:"permission_ids"`
}

type RoleAddPermissionsOutput struct{}

type RoleDeletePermissionsInput struct {
	RoleId        int   `json:"role_id"`
	PermissionIds []int `json:"permission_ids"`
}

type RoleDeletePermissionOutput struct{}
