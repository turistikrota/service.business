package req

import "github.com/turistikrota/service.owner/src/app/command"

type OwnershipUserPermRemoveRequest struct {
	NickName   string
	UserName   string
	Permission string `json:"permission" validate:"required"`
}

func (r *OwnershipUserPermRemoveRequest) LoadDetail(detail *OwnerShipDetailUserRequest) *OwnershipUserPermRemoveRequest {
	r.NickName = detail.NickName
	r.UserName = detail.UserName
	return r
}

func (r *OwnershipUserPermRemoveRequest) ToCommand(userUUID string) command.OwnershipUserPermRemoveCommand {
	return command.OwnershipUserPermRemoveCommand{
		OwnerNickName:  r.NickName,
		UserName:       r.UserName,
		AccessUserUUID: userUUID,
		PermissionName: r.Permission,
	}
}
