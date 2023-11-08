package req

import (
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
)

type OwnershipRequest struct {
	NickName string `params:"nickName" validate:"required"`
}

type OwnerShipDetailRequest struct {
	OwnershipRequest
}

func (o *OwnershipRequest) ToViewQuery() query.ViewOwnershipQuery {
	return query.ViewOwnershipQuery{
		NickName: o.NickName,
	}
}

func (o *OwnershipRequest) ToVerifyCommand(adminUUID string) command.AdminOwnershipVerifyCommand {
	return command.AdminOwnershipVerifyCommand{
		OwnerNickName: o.NickName,
		AdminUUID:     adminUUID,
	}
}

func (o *OwnerShipDetailRequest) ToGetWithUserOwnershipQuery(userUUID string, userName string) query.GetWithUserOwnershipQuery {
	return query.GetWithUserOwnershipQuery{
		NickName: o.NickName,
		UserName: userName,
		UserUUID: userUUID,
	}
}

func (o *OwnerShipDetailRequest) ToDisableCommand(userUUID string, userName string) command.OwnershipDisableCommand {
	return command.OwnershipDisableCommand{
		OwnerNickName: o.NickName,
		UserName:      userName,
		UserUUID:      userUUID,
	}
}

func (o *OwnerShipDetailRequest) ToEnableCommand(userUUID string, userName string) command.OwnershipEnableCommand {
	return command.OwnershipEnableCommand{
		OwnerNickName: o.NickName,
		UserName:      userName,
		UserUUID:      userUUID,
	}
}

func (o *OwnerShipDetailRequest) ToUserListQuery(userName string) query.ListMyOwnershipUsersQuery {
	return query.ListMyOwnershipUsersQuery{
		NickName: o.NickName,
		UserName: userName,
	}
}
