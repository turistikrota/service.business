package req

import "github.com/turistikrota/service.owner/src/app/command"

type OwnershipUserPermAddRequest struct {
	NickName   string
	UserName   string
	UserCode   string
	Permission string `json:"permission" validate:"required"`
}

func (r *OwnershipUserPermAddRequest) LoadDetail(detail *OwnerShipDetailUserRequest) *OwnershipUserPermAddRequest {
	detail.Parse()
	r.NickName = detail.NickName
	r.UserName = detail.UserName
	r.UserCode = detail.UserCode
	return r
}

func (r *OwnershipUserPermAddRequest) ToCommand(userUUID string) command.OwnershipUserPermAddCommand {
	return command.OwnershipUserPermAddCommand{
		OwnerNickName:  r.NickName,
		UserName:       r.UserName,
		UserCode:       r.UserCode,
		AccessUserUUID: userUUID,
		PermissionName: r.Permission,
	}
}
