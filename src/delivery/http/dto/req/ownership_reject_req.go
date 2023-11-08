package req

import "github.com/turistikrota/service.owner/src/app/command"

type OwnershipRejectRequest struct {
	Reason string `json:"reason" validate:"required"`
}

func (o *OwnershipRejectRequest) ToCommand(nickName string, adminUUID string) command.AdminOwnershipRejectCommand {
	return command.AdminOwnershipRejectCommand{
		OwnerNickName: nickName,
		AdminUUID:     adminUUID,
		Reason:        o.Reason,
	}
}
