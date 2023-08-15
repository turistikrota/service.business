package dto

import (
	"time"

	"github.com/turistikrota/service.owner/src/app/command"
)

type AccountCreated struct {
	UserUUID    string     `json:"user_uuid"`
	AccountName string     `json:"name"`
	CreatedAt   *time.Time `json:"created_at"`
}

func (e *AccountCreated) ToCommand() command.AccountCreateCommand {
	return command.AccountCreateCommand{
		UserUUID:    e.UserUUID,
		AccountName: e.AccountName,
		CreatedAt:   e.CreatedAt,
	}
}
