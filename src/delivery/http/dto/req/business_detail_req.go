package req

import (
	"github.com/turistikrota/service.business/src/app/command"
	"github.com/turistikrota/service.business/src/app/query"
)

type BusinessRequest struct {
	NickName string `params:"nickName" validate:"required"`
}

type BusinessShipDetailRequest struct {
	BusinessRequest
}

func (o *BusinessRequest) ToViewQuery() query.ViewBusinessQuery {
	return query.ViewBusinessQuery{
		NickName: o.NickName,
	}
}

func (o *BusinessRequest) ToVerifyCommand(adminUUID string) command.AdminBusinessVerifyCommand {
	return command.AdminBusinessVerifyCommand{
		BusinessNickName: o.NickName,
		AdminUUID:        adminUUID,
	}
}

func (o *BusinessRequest) ToDeleteCommand(adminUUID string) command.AdminBusinessDeleteCommand {
	return command.AdminBusinessDeleteCommand{
		BusinessNickName: o.NickName,
		AdminUUID:        adminUUID,
	}
}

func (o *BusinessRequest) ToRecoverCommand(adminUUID string) command.AdminBusinessRecoverCommand {
	return command.AdminBusinessRecoverCommand{
		BusinessNickName: o.NickName,
		AdminUUID:        adminUUID,
	}
}

func (o *BusinessRequest) ToAdminViewQuery() query.AdminViewBusinessQuery {
	return query.AdminViewBusinessQuery{
		NickName: o.NickName,
	}
}

func (o *BusinessShipDetailRequest) ToGetWithUserBusinessQuery(userUUID string, userName string) query.GetWithUserBusinessQuery {
	return query.GetWithUserBusinessQuery{
		NickName: o.NickName,
		UserName: userName,
		UserUUID: userUUID,
	}
}

func (o *BusinessShipDetailRequest) ToDisableCommand(userUUID string, userName string) command.BusinessDisableCommand {
	return command.BusinessDisableCommand{
		BusinessNickName: o.NickName,
		UserName:         userName,
		UserUUID:         userUUID,
	}
}

func (o *BusinessShipDetailRequest) ToEnableCommand(userUUID string, userName string) command.BusinessEnableCommand {
	return command.BusinessEnableCommand{
		BusinessNickName: o.NickName,
		UserName:         userName,
		UserUUID:         userUUID,
	}
}

func (o *BusinessShipDetailRequest) ToUserListQuery(userName string) query.ListMyBusinessUsersQuery {
	return query.ListMyBusinessUsersQuery{
		NickName: o.NickName,
		UserName: userName,
	}
}
