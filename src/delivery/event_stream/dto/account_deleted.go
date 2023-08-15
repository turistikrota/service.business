package dto

import "github.com/turistikrota/service.owner/src/app/command"

type AccountDeleted struct {
	UserUUID    string `json:"user_uuid"`
	AccountName string `json:"name"`
	AccountCode string `json:"code"`
}

func (e *AccountDeleted) ToCommand() command.AccountDeleteCommand {
	return command.AccountDeleteCommand{
		UserUUID: e.UserUUID,
		Name:     e.AccountName,
		Code:     e.AccountCode,
	}
}
