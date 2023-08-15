package dto

import "github.com/turistikrota/service.owner/src/app/command"

type AccountDisabled struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
}

func (e *AccountDisabled) ToCommand() command.AccountDisableCommand {
	return command.AccountDisableCommand{
		UserUUID: e.UserUUID,
		Name:     e.AccountName,
	}
}
