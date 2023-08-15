package req

import (
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
)

type OwnershipRequest struct {
	NickName string `params:"nickName" validate:"required"`
}

type OwnerShipDetailRequest struct {
	AccountUserRequest
	OwnershipRequest
}

func (o *OwnershipRequest) ToViewQuery() query.ViewOwnershipQuery {
	return query.ViewOwnershipQuery{
		NickName: o.NickName,
	}
}

func (o *OwnerShipDetailRequest) ToGetWithUserOwnershipQuery(userUUID string) query.GetWithUserOwnershipQuery {
	o.Parse()
	return query.GetWithUserOwnershipQuery{
		NickName: o.NickName,
		UserName: o.UserName,
		UserCode: o.UserCode,
		UserUUID: userUUID,
	}
}

func (o *OwnerShipDetailRequest) ToDisableCommand(userUUID string) command.OwnershipDisableCommand {
	o.Parse()
	return command.OwnershipDisableCommand{
		OwnerNickName: o.NickName,
		UserName:      o.UserName,
		UserCode:      o.UserCode,
		UserUUID:      userUUID,
	}
}

func (o *OwnerShipDetailRequest) ToEnableCommand(userUUID string) command.OwnershipEnableCommand {
	o.Parse()
	return command.OwnershipEnableCommand{
		OwnerNickName: o.NickName,
		UserName:      o.UserName,
		UserCode:      o.UserCode,
		UserUUID:      userUUID,
	}
}

func (o *OwnerShipDetailRequest) ToUserListQuery() query.ListMyOwnershipUsersQuery {
	o.Parse()
	return query.ListMyOwnershipUsersQuery{
		NickName: o.NickName,
		UserName: o.UserName,
		UserCode: o.UserCode,
	}
}
