package req

import (
	"time"

	"github.com/mixarchitecture/microp/formats"
	"github.com/turistikrota/service.business/src/app/command"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessApplicationRequest struct {
	NickName        string `json:"nickName" validate:"required"`
	RealName        string `json:"realName" validate:"required"`
	BusinessType    string `json:"businessType" validate:"required,oneof=individual corporation"`
	FirstName       string `json:"firstName" validate:"required_if=BusinessType individual"`
	LastName        string `json:"lastName" validate:"required_if=BusinessType individual"`
	IdentityNumber  string `json:"identityNumber" validate:"required_if=BusinessType individual"`
	SerialNumber    string `json:"serialNumber" validate:"required_if=BusinessType individual"`
	Province        string `json:"province" validate:"required"`
	District        string `json:"district" validate:"required"`
	Address         string `json:"address" validate:"required"`
	DateOfBirth     string `json:"dateOfBirth" validate:"required_if=BusinessType individual,datetime=2006-01-02"`
	TaxNumber       string `json:"taxNumber" validate:"required_if=BusinessType corporation"`
	CorporationType string `json:"type" validate:"required_if=BusinessType corporation"`
}

func (r *BusinessApplicationRequest) ToCommand(userUUID string, userName string) command.BusinessApplicationCommand {
	businessType := business.Type(r.BusinessType)
	cmd := command.BusinessApplicationCommand{
		UserName:     userName,
		UserUUID:     userUUID,
		NickName:     r.NickName,
		RealName:     r.RealName,
		BusinessType: businessType,
	}
	if businessType == business.Types.Individual {
		birth, _ := time.Parse(formats.DateYYYYMMDD, r.DateOfBirth)
		cmd.Individual = command.BusinessApplicationIndividualCommand{
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
		cmd.Corporation = command.BusinessApplicationCorporationCommand{
			TaxNumber: r.TaxNumber,
			Province:  r.Province,
			District:  r.District,
			Address:   r.Address,
			Type:      r.CorporationType,
		}
	}
	return cmd
}
