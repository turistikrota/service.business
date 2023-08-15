package req

import (
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.shared/helper"
)

type OwnerShipDetailUserRequest struct {
	NickName        string `params:"nickName" validate:"required"`
	UserNameAndCode string `params:"userName" validate:"required,username_and_code"`
	UserName        string
	UserCode        string
}

func (r *OwnerShipDetailUserRequest) Parse() *OwnerShipDetailUserRequest {
	_, r.UserName, r.UserCode = helper.Parsers.ParseUsernameAndCode(r.UserNameAndCode)
	return r
}

func (o *OwnerShipDetailUserRequest) ToGetWithUserOwnershipQuery() query.GetWithUserOwnershipQuery {
	o.Parse()
	return query.GetWithUserOwnershipQuery{
		NickName: o.NickName,
		UserName: o.UserName,
		UserCode: o.UserCode,
	}
}

func (o *OwnerShipDetailUserRequest) ToViewQuery() query.ViewOwnershipQuery {
	return query.ViewOwnershipQuery{
		NickName: o.NickName,
	}
}

func (o *OwnerShipDetailUserRequest) ToAddUserCommand(userUUID string) command.OwnershipUserAddCommand {
	o.Parse()
	return command.OwnershipUserAddCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
		UserCode:       o.UserCode,
	}
}

func (o *OwnerShipDetailUserRequest) ToAddUserPermCommand(userUUID string) command.OwnershipUserPermAddCommand {
	o.Parse()
	return command.OwnershipUserPermAddCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
		UserCode:       o.UserCode,
	}
}

func (o *OwnerShipDetailUserRequest) ToRemoveUserPermCommand(userUUID string) command.OwnershipUserPermRemoveCommand {
	o.Parse()
	return command.OwnershipUserPermRemoveCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
		UserCode:       o.UserCode,
	}
}

func (o *OwnerShipDetailUserRequest) ToRemoveUserCommand(userUUID string) command.OwnershipUserRemoveCommand {
	o.Parse()
	return command.OwnershipUserRemoveCommand{
		OwnerNickName:  o.NickName,
		AccessUserUUID: userUUID,
		UserName:       o.UserName,
		UserCode:       o.UserCode,
	}
}
