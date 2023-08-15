package owner

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
	CorporationTypeRequired() *i18np.Error
	CorporationTypeInvalid() *i18np.Error
	IdentityVerificationFailed() *i18np.Error
}

type ownerErrors struct{}

func newOwnerErrors() Errors {
	return &ownerErrors{}
}

func (e *ownerErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *ownerErrors) NickNameAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.NickNameAlreadyExists)
}

func (e *ownerErrors) IdentityNumberAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.IdentityNumberAlreadyExists)
}

func (e *ownerErrors) TaxNumberAlreadyExists() *i18np.Error {
	return i18np.NewError(I18nMessages.TaxNumberAlreadyExists)
}

func (e *ownerErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound)
}

func (e *ownerErrors) TypeRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.TypeRequired)
}

func (e *ownerErrors) TypeInvalid() *i18np.Error {
	return i18np.NewError(I18nMessages.TypeInvalid)
}

func (e *ownerErrors) CorporationTypeRequired() *i18np.Error {
	return i18np.NewError(I18nMessages.CorporationTypeRequired)
}

func (e *ownerErrors) CorporationTypeInvalid() *i18np.Error {
	return i18np.NewError(I18nMessages.CorporationTypeInvalid)
}

func (e *ownerErrors) IdentityVerificationFailed() *i18np.Error {
	return i18np.NewError(I18nMessages.IdentityVerificationFailed)
}
