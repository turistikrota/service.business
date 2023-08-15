package dto

type Dto interface {
	AccountCreated() *AccountCreated
	AccountDisabled() *AccountDisabled
	AccountEnabled() *AccountEnabled
	AccountDeleted() *AccountDeleted
	AccountUpdated() *AccountUpdated
}

type dto struct{}

func New() Dto {
	return &dto{}
}

func (d *dto) AccountCreated() *AccountCreated {
	return &AccountCreated{}
}

func (d *dto) AccountDisabled() *AccountDisabled {
	return &AccountDisabled{}
}

func (d *dto) AccountEnabled() *AccountEnabled {
	return &AccountEnabled{}
}

func (d *dto) AccountDeleted() *AccountDeleted {
	return &AccountDeleted{}
}

func (d *dto) AccountUpdated() *AccountUpdated {
	return &AccountUpdated{}
}
