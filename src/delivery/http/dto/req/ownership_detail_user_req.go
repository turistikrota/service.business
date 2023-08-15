package req

import (
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
)

type OwnerShipDetailUserRequest struct {
	NickName string `params:"nickName" validate:"required"`
	UserName string `params:"userName" validate:"required,username"`
}

func (o *OwnerShipDetailUserRequest) ToGetWithUserOwnershipQuery() query.GetWithUserOwnershipQuery {
	return query.GetWithUserOwnershipQuery{
		NickName: o.NickName,
		UserName: o.UserName,
	}
}

func (o *OwnerShipDetailUserRequest) ToViewQuery() query.ViewOwnershipQuery {
	return query.ViewOwnershipQuery{
		NickName: o.NickName,
	}
}

func (o *OwnerShipDetailUserRequest) ToAddUserCommand(userUUID string) command.OwnershipUserAddCommand {
	return command.OwnershipUserAddCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
	}
}

func (o *OwnerShipDetailUserRequest) ToAddUserPermCommand(userUUID string) command.OwnershipUserPermAddCommand {
	return command.OwnershipUserPermAddCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
	}
}

func (o *OwnerShipDetailUserRequest) ToRemoveUserPermCommand(userUUID string) command.OwnershipUserPermRemoveCommand {
	return command.OwnershipUserPermRemoveCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
	}
}

func (o *OwnerShipDetailUserRequest) ToRemoveUserCommand(userUUID string) command.OwnershipUserRemoveCommand {
	return command.OwnershipUserRemoveCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
	}
}
