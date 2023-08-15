package dto

import "github.com/turistikrota/service.owner/src/app/command"

type AccountEnabled struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
}

func (e *AccountEnabled) ToCommand() command.AccountEnableCommand {
	return command.AccountEnableCommand{
		UserUUID: e.UserUUID,
		Name:     e.AccountName,
	}
}
