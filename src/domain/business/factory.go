package business

import (
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.business/src/config"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{Errors: newBusinessErrors()}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewBusinessParams struct {
	UserUUID     string
	UserName     string
	UserCode     string
	NickName     string
	RealName     string
	BusinessType Type
	Individual   Individual
	Corporation  Corporation
}

func (f Factory) NewUser(uuid string, name string) *User {
	return &User{
		UUID: uuid,
		Name: name,
		Roles: []string{
			config.Roles.Business.Member,
		},
		JoinAt: time.Now(),
	}
}

func (f Factory) NewBusiness(params NewBusinessParams) *Entity {
	t := time.Now()
	e := &Entity{
		NickName:     params.NickName,
		RealName:     params.RealName,
		BusinessType: params.BusinessType,
		Users: []User{
			{
				UUID: params.UserUUID,
				Name: params.UserName,
				Roles: []string{
					config.Roles.Business.Super,
					config.Roles.Business.AdminView,
					config.Roles.Business.Enable,
					config.Roles.Business.Disable,
					config.Roles.Business.UserAdd,
					config.Roles.Business.UserRemove,
					config.Roles.Business.UserList,
					config.Roles.Business.UserPermRemove,
					config.Roles.Business.UserPermAdd,
					config.Roles.Business.Member,
					config.Roles.Business.InviteCreate,
					config.Roles.Business.InviteDelete,
					config.Roles.Business.InviteView,
					config.Roles.Business.UploadAvatar,
					config.Roles.Business.UploadCover,
				},
				JoinAt: t,
			},
		},
		IsEnabled:  false,
		IsVerified: false,
		VerifiedAt: nil,
		CreatedAt:  &t,
		UpdatedAt:  &t,
	}
	if params.BusinessType == Types.Individual {
		e.Individual = params.Individual
	} else {
		e.Corporation = params.Corporation
	}
	return e
}

func (f Factory) Validate(e *Entity) *i18np.Error {
	if err := f.validateType(e); err != nil {
		return err
	}
	return f.validateByType(e)
}

func (f Factory) validateType(e *Entity) *i18np.Error {
	if e.BusinessType == "" {
		return f.Errors.TypeRequired()
	}
	if e.BusinessType != Types.Individual && e.BusinessType != Types.Corporation {
		return f.Errors.TypeInvalid()
	}
	return nil
}

func (f Factory) validateByType(e *Entity) *i18np.Error {
	if e.BusinessType == Types.Individual {
		return f.validateIndividual(e)
	}
	return f.validateCorporation(e)
}

func (f Factory) validateIndividual(e *Entity) *i18np.Error {
	return nil
}

func (f Factory) validateCorporation(e *Entity) *i18np.Error {
	if e.Corporation.Type == "" {
		return f.Errors.CorporationTypeRequired()
	}
	switch e.Corporation.Type {
	case CorporationTypes.Anonymous,
		CorporationTypes.SoleProprietorship,
		CorporationTypes.Limited,
		CorporationTypes.Collective,
		CorporationTypes.Cooperative,
		CorporationTypes.OrdinaryPartnership,
		CorporationTypes.OrdinaryLimitedPartnership,
		CorporationTypes.LimitedPartnershipShareCapital,
		CorporationTypes.Other:
		return nil
	}
	return f.Errors.CorporationTypeInvalid()
}
