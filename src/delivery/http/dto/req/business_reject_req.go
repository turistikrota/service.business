package req

import "github.com/turistikrota/service.business/src/app/command"

type BusinessRejectRequest struct {
	Reason string `json:"reason" validate:"required"`
}

func (o *BusinessRejectRequest) ToCommand(nickName string, adminUUID string) command.AdminBusinessRejectCommand {
	return command.AdminBusinessRejectCommand{
		BusinessNickName: nickName,
		AdminUUID:        adminUUID,
		Reason:           o.Reason,
	}
}
