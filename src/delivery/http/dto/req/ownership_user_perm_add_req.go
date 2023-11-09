package req

import "github.com/turistikrota/service.owner/src/app/command"

type OwnershipUserPermAddRequest struct {
	NickName   string
	UserName   string
	Permission string `json:"permission" validate:"required"`
}

func (r *OwnershipUserPermAddRequest) LoadDetail(detail *OwnerShipDetailUserRequest) *OwnershipUserPermAddRequest {
	r.NickName = detail.NickName
	r.UserName = detail.UserName
	return r
}

func (r *OwnershipUserPermAddRequest) ToCommand(userUUID string, userName string) command.OwnershipUserPermAddCommand {
	return command.OwnershipUserPermAddCommand{
		OwnerNickName:  r.NickName,
		UserName:       r.UserName,
		AccessUserUUID: userUUID,
		AccessUserName: userName,
		PermissionName: r.Permission,
	}
}
