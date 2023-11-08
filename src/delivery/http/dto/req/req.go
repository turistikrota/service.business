package req

type Request interface {
	OwnerShipDetail() *OwnerShipDetailRequest
	Ownership() *OwnershipRequest
	OwnerApplication() *OwnerApplicationRequest
	OwnerShipDetailUser() *OwnerShipDetailUserRequest
	OwnerPermissionAdd() *OwnershipUserPermAddRequest
	OwnerPermissionRemove() *OwnershipUserPermRemoveRequest
	OwnerSelect() *OwnershipSelectRequest
	Pagination() *PaginationRequest
	InviteCreate() *InviteCreateRequest
	InviteDetail() *InviteDetailRequest
	OwnershipReject() *OwnershipRejectRequest
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

func (r *request) Ownership() *OwnershipRequest {
	return &OwnershipRequest{}
}

func (r *request) OwnerPermissionAdd() *OwnershipUserPermAddRequest {
	return &OwnershipUserPermAddRequest{}
}

func (r *request) OwnerPermissionRemove() *OwnershipUserPermRemoveRequest {
	return &OwnershipUserPermRemoveRequest{}
}

func (r *request) OwnerSelect() *OwnershipSelectRequest {
	return &OwnershipSelectRequest{}
}

func (r *request) InviteCreate() *InviteCreateRequest {
	return &InviteCreateRequest{}
}

func (r *request) InviteDetail() *InviteDetailRequest {
	return &InviteDetailRequest{}
}

func (r *request) OwnershipReject() *OwnershipRejectRequest {
	return &OwnershipRejectRequest{}
}
