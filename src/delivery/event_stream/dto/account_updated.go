package dto

import (
	"time"

	"github.com/turistikrota/service.owner/src/app/command"
)

type AccountUpdated struct {
	UserUUID    string              `json:"user_uuid"`
	AccountName string              `json:"name"`
	AccountCode string              `json:"code"`
	Entity      AccountUpdateEntity `json:"entity"`
}

type AccountUpdateEntity struct {
	UserName  string     `json:"user_name"`
	UserCode  string     `json:"user_code"`
	FullName  string     `json:"full_name"`
	AvatarURL string     `json:"avatar_url"`
	BirthDate *time.Time `json:"birth_date"`
}

func (e *AccountUpdated) ToCommand() command.AccountUpdateCommand {
	return command.AccountUpdateCommand{
		UserUUID:    e.UserUUID,
		CurrentName: e.AccountName,
		CurrentCode: e.AccountCode,
		NewName:     e.Entity.UserName,
		NewCode:     e.Entity.UserCode,
		FullName:    e.Entity.FullName,
		Avatar:      e.Entity.AvatarURL,
		BirthDate:   e.Entity.BirthDate,
	}
}
