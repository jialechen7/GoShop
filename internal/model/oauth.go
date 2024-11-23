package model

type GithubReceiveCodeInput struct {
	Code  string
	State string
}

type GithubReceiveCodeOutput struct {
	Type        string      `json:"type"`
	Token       string      `json:"token"`
	ExpireIn    int         `json:"expire_in"`
	IsAdmin     int         `json:"is_admin"`
	RoleIds     string      `json:"role_ids"`
	Permissions interface{} `json:"permissions"`
}
