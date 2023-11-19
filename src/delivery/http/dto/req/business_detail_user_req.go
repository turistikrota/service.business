package req

import (
	"github.com/turistikrota/service.business/src/app/command"
	"github.com/turistikrota/service.business/src/app/query"
)

type BusinessShipDetailUserRequest struct {
	NickName string `params:"nickName" validate:"required"`
	UserName string `params:"userName" validate:"required,username"`
}

func (o *BusinessShipDetailUserRequest) ToGetWithUserBusinessQuery() query.GetWithUserBusinessQuery {
	return query.GetWithUserBusinessQuery{
		NickName: o.NickName,
		UserName: o.UserName,
	}
}

func (o *BusinessShipDetailUserRequest) ToViewQuery() query.ViewBusinessQuery {
	return query.ViewBusinessQuery{
		NickName: o.NickName,
	}
}

func (o *BusinessShipDetailUserRequest) ToAddUserPermCommand(userUUID string) command.BusinessUserPermAddCommand {
	return command.BusinessUserPermAddCommand{
		BusinessNickName: o.NickName,
		AccessUserUUID:   userUUID,
		UserName:         o.UserName,
	}
}

func (o *BusinessShipDetailUserRequest) ToRemoveUserPermCommand(userUUID string) command.BusinessUserPermRemoveCommand {
	return command.BusinessUserPermRemoveCommand{
		BusinessNickName: o.NickName,
		AccessUserUUID:   userUUID,
		UserName:         o.UserName,
	}
}

func (o *BusinessShipDetailUserRequest) ToRemoveUserCommand(userUUID string, userName string) command.BusinessUserRemoveCommand {
	return command.BusinessUserRemoveCommand{
		BusinessNickName: o.NickName,
		AccessUserUUID:   userUUID,
		AccessUserName:   userName,
		UserName:         o.UserName,
	}
}
