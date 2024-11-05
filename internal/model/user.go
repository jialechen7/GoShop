package model

type UserCreateUpdateBase struct {
	Name         string
	Password     string
	UserSalt     string
	Avatar       string
	Sign         string
	Sex          int
	SecretAnswer string
}

type UserCreateInput struct {
	UserCreateUpdateBase
}

type UserCreateOutput struct {
	Id int `json:"id"`
}

type UserUpdateInput struct {
	UserCreateUpdateBase
	Id int
}

type UserUpdateOutput struct{}
