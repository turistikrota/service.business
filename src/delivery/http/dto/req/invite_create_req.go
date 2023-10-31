package req

import "github.com/turistikrota/service.owner/src/app/command"

type InviteCreateRequest struct {
	Email  string `json:"email" validate:"required,email"`
	Locale string `json:"locale" validate:"required,locale"`
}

func (r *InviteCreateRequest) ToCommand(nickName string, userName string) command.InviteCreateCommand {
	return command.InviteCreateCommand{
		Email:  r.Email,
		Locale: r.Locale,
		CreatorUserName: userName,
		OwnerNickName:   nickName,
	}
}