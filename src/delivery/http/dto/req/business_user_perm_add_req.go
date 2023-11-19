package req

import "github.com/turistikrota/service.business/src/app/command"

type BusinessUserPermAddRequest struct {
	NickName   string
	UserName   string
	Permission string `json:"permission" validate:"required"`
}

func (r *BusinessUserPermAddRequest) LoadDetail(detail *BusinessShipDetailUserRequest) *BusinessUserPermAddRequest {
	r.NickName = detail.NickName
	r.UserName = detail.UserName
	return r
}

func (r *BusinessUserPermAddRequest) ToCommand(userUUID string, userName string) command.BusinessUserPermAddCommand {
	return command.BusinessUserPermAddCommand{
		BusinessNickName: r.NickName,
		UserName:         r.UserName,
		AccessUserUUID:   userUUID,
		AccessUserName:   userName,
		PermissionName:   r.Permission,
	}
}
