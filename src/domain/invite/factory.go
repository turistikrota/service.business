package invite

import "time"

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{Errors: newInviteErrors()}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

func (f Factory) New(email string, ownerUUID string) *Entity {
	t := time.Now()
	return &Entity{
		Email:     email,
		OwnerUUID: ownerUUID,
		IsUsed:    false,
		IsDeleted: false,
		CreatedAt: &t,
		UpdatedAt: nil,
	}
}
