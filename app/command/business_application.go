package command

import (
	"context"
	"time"

	"github.com/9ssi7/vkn"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/formats"
	"github.com/cilloparch/cillop/i18np"
	"github.com/ssibrahimbas/KPSPublic"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.shared/cipher"
)

type BusinessApplicationCmd struct {
	UserUUID        string
	UserName        string
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
	Application     string `json:"application" validate:"required,oneof=accommodation advert place"`
	CorporationType string `json:"type" validate:"required_if=BusinessType corporation"`
	Locale          business.Locale
}

type BusinessApplicationRes struct {
	UUID string `json:"uuid"`
}

type BusinessApplicationHandler cqrs.HandlerFunc[BusinessApplicationCmd, *BusinessApplicationRes]

func NewBusinessApplicationHandler(repo business.Repository, factory business.Factory, events business.Events, identitySrv KPSPublic.Service, taxIdSrv vkn.Vkn, cipher cipher.Service) BusinessApplicationHandler {
	return func(ctx context.Context, cmd BusinessApplicationCmd) (*BusinessApplicationRes, *i18np.Error) {
		params := business.NewBusinessParams{
			UserUUID:     cmd.UserUUID,
			UserName:     cmd.UserName,
			NickName:     cmd.NickName,
			RealName:     cmd.RealName,
			BusinessType: business.Type(cmd.BusinessType),
			Application:  business.Application(cmd.Application),
			Locale:       cmd.Locale,
		}
		if params.BusinessType == business.Types.Individual {
			birth, _ := time.Parse(formats.DateYYYYMMDD, cmd.DateOfBirth)
			res, err := identitySrv.Verify(KPSPublic.VerifyConfig{
				SerialNumber:     cmd.SerialNumber,
				TCIdentityNumber: cmd.IdentityNumber,
				LastName:         cmd.LastName,
				FirstName:        cmd.FirstName,
				BirthYear:        birth.Year(),
				BirthMonth:       int(birth.Month()),
				BirthDay:         birth.Day(),
			})
			if err != nil {
				return nil, factory.Errors.Failed(err.Error())
			}
			if !res {
				return nil, factory.Errors.IdentityVerificationFailed()
			}
			identity, error := cipher.Encrypt(cmd.IdentityNumber)
			if error != nil {
				return nil, factory.Errors.Failed("failed to hash identity number")
			}
			serial, error := cipher.Encrypt(cmd.SerialNumber)
			if error != nil {
				return nil, factory.Errors.Failed("failed to hash serial number")
			}
			params.Individual = business.Individual{
				FirstName:      cmd.FirstName,
				LastName:       cmd.LastName,
				IdentityNumber: identity,
				SerialNumber:   serial,
				Province:       cmd.Province,
				District:       cmd.District,
				Address:        cmd.Address,
				DateOfBirth:    birth,
			}
			_, notFound, error := repo.GetByIndividual(ctx, params.Individual)
			if error != nil {
				return nil, error
			}
			if !notFound {
				return nil, factory.Errors.IndividualAlreadyExists()
			}
		} else {
			res, _err := taxIdSrv.GetRecipient(cmd.TaxNumber)
			if _err != nil {
				return nil, factory.Errors.Failed(_err.Error())
			}
			if res == nil || res.Data.TaxOffice == "" || res.Data.Title == "" {
				return nil, factory.Errors.CorporationVerificationFailed()
			}
			tax, err := cipher.Encrypt(cmd.TaxNumber)
			if err != nil {
				return nil, factory.Errors.Failed("failed to hash tax number")
			}
			params.Corporation = business.Corporation{
				TaxNumber: tax,
				Province:  cmd.Province,
				District:  cmd.District,
				Address:   cmd.Address,
				TaxOffice: res.Data.TaxOffice,
				Title:     res.Data.Title,
				Type:      business.CorporationType(cmd.CorporationType),
			}
			_, notFound, err := repo.GetByCorporation(ctx, params.Corporation)
			if err != nil {
				return nil, err
			}
			if !notFound {
				return nil, factory.Errors.CorporationAlreadyExists()
			}
		}
		entity := factory.NewBusiness(params)
		if err := factory.Validate(entity); err != nil {
			return nil, err
		}
		res, err := repo.CheckNickName(ctx, entity.NickName)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		if res {
			return nil, factory.Errors.NickNameAlreadyExists()
		}
		inserted, err := repo.Create(ctx, entity)
		if err != nil {
			return nil, factory.Errors.Failed("failed to save business")
		}
		events.Created(&business.EventBusinessCreated{
			Business: inserted,
			UserUUID: cmd.UserUUID,
			UserName: cmd.UserName,
		})
		return &BusinessApplicationRes{
			UUID: inserted.UUID,
		}, nil
	}
}
