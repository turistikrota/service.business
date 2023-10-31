package owner

import (
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/config"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{Errors: newOwnerErrors()}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewOwnerParams struct {
	UserUUID    string
	UserName    string
	UserCode    string
	NickName    string
	RealName    string
	OwnerType   Type
	Individual  Individual
	Corporation Corporation
}

func (f Factory) NewUser(uuid string, name string) *User {
	return &User{
		UUID: uuid,
		Name: name,
		Roles: []string{
			config.Roles.Owner.Member,
		},
		JoinAt: time.Now(),
	}
}

func (f Factory) NewOwner(params NewOwnerParams) *Entity {
	t := time.Now()
	e := &Entity{
		NickName:  params.NickName,
		RealName:  params.RealName,
		AvatarURL: "",
		CoverURL:  "",
		OwnerType: params.OwnerType,
		Users: []User{
			{
				UUID: params.UserUUID,
				Name: params.UserName,
				Roles: []string{
					config.Roles.Owner.Super,
					config.Roles.Owner.AdminView,
					config.Roles.Owner.Enable,
					config.Roles.Owner.Disable,
					config.Roles.Owner.UserAdd,
					config.Roles.Owner.UserRemove,
					config.Roles.Owner.UserList,
					config.Roles.Owner.UserPermRemove,
					config.Roles.Owner.UserPermAdd,
					config.Roles.Owner.Member,
					config.Roles.Owner.InviteCreate,
					config.Roles.Owner.InviteDelete,
					config.Roles.Owner.InviteView,
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
	if params.OwnerType == Types.Individual {
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
	if e.OwnerType == "" {
		return f.Errors.TypeRequired()
	}
	if e.OwnerType != Types.Individual && e.OwnerType != Types.Corporation {
		return f.Errors.TypeInvalid()
	}
	return nil
}

func (f Factory) validateByType(e *Entity) *i18np.Error {
	if e.OwnerType == Types.Individual {
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
