package business

import "github.com/cilloparch/cillop/i18np"

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
	return i18np.NewError(Messages.Failed, i18np.P{"Operation": operation})
}

func (e *businessErrors) NickNameAlreadyExists() *i18np.Error {
	return i18np.NewError(Messages.NickNameAlreadyExists)
}

func (e *businessErrors) IdentityNumberAlreadyExists() *i18np.Error {
	return i18np.NewError(Messages.IdentityNumberAlreadyExists)
}

func (e *businessErrors) TaxNumberAlreadyExists() *i18np.Error {
	return i18np.NewError(Messages.TaxNumberAlreadyExists)
}

func (e *businessErrors) NotFound() *i18np.Error {
	return i18np.NewError(Messages.NotFound)
}

func (e *businessErrors) TypeRequired() *i18np.Error {
	return i18np.NewError(Messages.TypeRequired)
}

func (e *businessErrors) TypeInvalid() *i18np.Error {
	return i18np.NewError(Messages.TypeInvalid)
}

func (e *businessErrors) CorporationTypeRequired() *i18np.Error {
	return i18np.NewError(Messages.CorporationTypeRequired)
}

func (e *businessErrors) CorporationTypeInvalid() *i18np.Error {
	return i18np.NewError(Messages.CorporationTypeInvalid)
}

func (e *businessErrors) IdentityVerificationFailed() *i18np.Error {
	return i18np.NewError(Messages.IdentityVerificationFailed)
}

func (e *businessErrors) CorporationVerificationFailed() *i18np.Error {
	return i18np.NewError(Messages.CorporationVerificationFailed)
}

func (e *businessErrors) IndividualAlreadyExists() *i18np.Error {
	return i18np.NewError(Messages.IndividualAlreadyExists)
}

func (e *businessErrors) CorporationAlreadyExists() *i18np.Error {
	return i18np.NewError(Messages.CorporationAlreadyExists)
}

func (e *businessErrors) ApplicationInvalid() *i18np.Error {
	return i18np.NewError(Messages.ApplicationInvalid)
}

func (e *businessErrors) BusinessMustBeCorporation() *i18np.Error {
	return i18np.NewError(Messages.BusinessMustBeCorporation)
}
