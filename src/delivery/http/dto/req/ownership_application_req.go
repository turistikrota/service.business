package req

import (
	"time"

	"github.com/mixarchitecture/microp/formats"
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type OwnerApplicationRequest struct {
	NickName        string `json:"nickName" validate:"required"`
	RealName        string `json:"realName" validate:"required"`
	OwnerType       string `json:"ownerType" validate:"required,oneof=individual corporation"`
	FirstName       string `json:"firstName" validate:"required_if=OwnerType individual"`
	LastName        string `json:"lastName" validate:"required_if=OwnerType individual"`
	IdentityNumber  string `json:"identityNumber" validate:"required_if=OwnerType individual"`
	SerialNumber    string `json:"serialNumber" validate:"required_if=OwnerType individual"`
	Province        string `json:"province" validate:"required"`
	District        string `json:"district" validate:"required"`
	Address         string `json:"address" validate:"required"`
	DateOfBirth     string `json:"dateOfBirth" validate:"required_if=OwnerType individual,datetime=2006-01-02"`
	TaxNumber       string `json:"taxNumber" validate:"required_if=OwnerType corporation"`
	CorporationType string `json:"type" validate:"required_if=OwnerType corporation"`
}

func (r *OwnerApplicationRequest) ToCommand(userUUID string, userName string) command.OwnerApplicationCommand {
	ownerType := owner.Type(r.OwnerType)
	cmd := command.OwnerApplicationCommand{
		UserName:  userName,
		UserUUID:  userUUID,
		NickName:  r.NickName,
		RealName:  r.RealName,
		OwnerType: ownerType,
	}
	if ownerType == owner.Types.Individual {
		birth, _ := time.Parse(formats.DateYYYYMMDD, r.DateOfBirth)
		cmd.Individual = command.OwnerApplicationIndividualCommand{
			FirstName:      r.FirstName,
			LastName:       r.LastName,
			IdentityNumber: r.IdentityNumber,
			SerialNumber:   r.SerialNumber,
			Province:       r.Province,
			District:       r.District,
			Address:        r.Address,
			DateOfBirth:    birth,
		}
	} else {
		cmd.Corporation = command.OwnerApplicationCorporationCommand{
			TaxNumber: r.TaxNumber,
			Province:  r.Province,
			District:  r.District,
			Address:   r.Address,
			Type:      r.CorporationType,
		}
	}
	return cmd
}
