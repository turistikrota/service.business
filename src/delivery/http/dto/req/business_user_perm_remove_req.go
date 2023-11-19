package req

import "github.com/turistikrota/service.business/src/app/command"

type BusinessUserPermRemoveRequest struct {
	NickName   string
	UserName   string
	Permission string `json:"permission" validate:"required"`
}

func (r *BusinessUserPermRemoveRequest) LoadDetail(detail *BusinessShipDetailUserRequest) *BusinessUserPermRemoveRequest {
	r.NickName = detail.NickName
	r.UserName = detail.UserName
	return r
}

func (r *BusinessUserPermRemoveRequest) ToCommand(userUUID string, userName string) command.BusinessUserPermRemoveCommand {
	return command.BusinessUserPermRemoveCommand{
		BusinessNickName: r.NickName,
		UserName:         r.UserName,
		AccessUserUUID:   userUUID,
		AccessUserName:   userName,
		PermissionName:   r.Permission,
	}
}
