package command

type UserDetailCmd struct {
	UserName string `params:"userName" validate:"required,username"`
}
