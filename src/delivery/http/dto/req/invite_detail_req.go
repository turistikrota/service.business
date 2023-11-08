package req

import (
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
)

type InviteDetailRequest struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

func (r *InviteDetailRequest) ToDelete(userUUID string, accountName string) command.InviteDeleteCommand {
	return command.InviteDeleteCommand{
		InviteUUID: r.UUID,
		UserUUID:   userUUID,
		UserName:   accountName,
	}
}

func (r *InviteDetailRequest) ToGet() query.InviteGetByUUIDQuery {
	return query.InviteGetByUUIDQuery{
		UUID: r.UUID,
	}
}

func (r *InviteDetailRequest) ToUse(uuid string, email string, name string) command.InviteUseCommand {
	return command.InviteUseCommand{
		InviteUUID: r.UUID,
		UserEmail:  email,
		UserUUID:   uuid,
		UserName:   name,
	}
}
