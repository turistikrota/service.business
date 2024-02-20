package business

import (
	"github.com/mixarchitecture/i18np"
)

type Errors interface {
	Failed(operation string) *i18np.Error
	NickNameAlreadyExists() *i18np.Error
	IdentityNumberAlreadyExists() *i18np.Error
	TaxNumberAlreadyExists() *i18np.Error
	NotFound() *i18np.Error
	TypeRequired() *i18np.Error
	TypeInvalid() *i18np.Error
	BusinessMustBeCorporation() *i18np.Error
	CorporationTypeRequired() *i18np.Error
	CorporationTypeInvalid() *i18np.Error
	IdentityVerificationFailed() *i18np.Error
	CorporationVerificationFailed() *i18np.Error
	IndividualAlreadyExists() *i18np.Error
	CorporationAlreadyExists() *i18np.Error
	ApplicationInvalid() *i18np.Error
}

type businessErrors struct{}

func newBusinessErrors() Errors {
	return &businessErrors{}
}

func (e *businessErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *businessErrors) NickNameAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.NickNameAlreadyExists)
}

func (e *businessErrors) IdentityNumberAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.IdentityNumberAlreadyExists)
}

func (e *businessErrors) TaxNumberAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.TaxNumberAlreadyExists)
}

func (e *businessErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound)
}

func (e *businessErrors) TypeRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.TypeRequired)
}

func (e *businessErrors) TypeInvalid() *i18np.Error {
	return i18np.NewError(I18nMessages.TypeInvalid)
}

func (e *businessErrors) CorporationTypeRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.CorporationTypeRequired)
}

func (e *businessErrors) CorporationTypeInvalid() *i18np.Error {
	return i18np.NewError(I18nMessages.CorporationTypeInvalid)
}

func (e *businessErrors) IdentityVerificationFailed() *i18np.Error {
	return i18np.NewError(I18nMessages.IdentityVerificationFailed)
}

func (e *businessErrors) CorporationVerificationFailed() *i18np.Error {
	return i18np.NewError(I18nMessages.CorporationVerificationFailed)
}

func (e *businessErrors) IndividualAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.IndividualAlreadyExists)
}

func (e *businessErrors) CorporationAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.CorporationAlreadyExists)
}

func (e *businessErrors) ApplicationInvalid() *i18np.Error {
	return i18np.NewError(I18nMessages.ApplicationInvalid)
}

func (e *businessErrors) BusinessMustBeCorporation() *i18np.Error {
	return i18np.NewError(I18nMessages.BusinessMustBeCorporation)
}
