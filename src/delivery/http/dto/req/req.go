package req

type Request interface {
	OwnerShipDetail() *OwnerShipDetailRequest
	Ownership() *OwnershipRequest
	OwnerApplication() *OwnerApplicationRequest
	OwnerShipDetailUser() *OwnerShipDetailUserRequest
	UserAccount() *AccountUserRequest
	OwnerPermissionAdd() *OwnershipUserPermAddRequest
	OwnerPermissionRemove() *OwnershipUserPermRemoveRequest
}

type request struct{}

func New() Request {
	return &request{}
}

func (r *request) OwnerShipDetail() *OwnerShipDetailRequest {
	return &OwnerShipDetailRequest{}
}

func (r *request) OwnerApplication() *OwnerApplicationRequest {
	return &OwnerApplicationRequest{}
}

func (r *request) OwnerShipDetailUser() *OwnerShipDetailUserRequest {
	return &OwnerShipDetailUserRequest{}
}

func (r *request) UserAccount() *AccountUserRequest {
	return &AccountUserRequest{}
}

func (r *request) Ownership() *OwnershipRequest {
	return &OwnershipRequest{}
}

func (r *request) OwnerPermissionAdd() *OwnershipUserPermAddRequest {
	return &OwnershipUserPermAddRequest{}
}

func (r *request) OwnerPermissionRemove() *OwnershipUserPermRemoveRequest {
	return &OwnershipUserPermRemoveRequest{}
}
