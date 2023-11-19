package req

type Request interface {
	BusinessShipDetail() *BusinessShipDetailRequest
	Business() *BusinessRequest
	BusinessApplication() *BusinessApplicationRequest
	BusinessShipDetailUser() *BusinessShipDetailUserRequest
	BusinessPermissionAdd() *BusinessUserPermAddRequest
	BusinessPermissionRemove() *BusinessUserPermRemoveRequest
	BusinessSelect() *BusinessSelectRequest
	Pagination() *PaginationRequest
	InviteCreate() *InviteCreateRequest
	InviteDetail() *InviteDetailRequest
	BusinessReject() *BusinessRejectRequest
}

type request struct{}

func New() Request {
	return &request{}
}

func (r *request) BusinessShipDetail() *BusinessShipDetailRequest {
	return &BusinessShipDetailRequest{}
}

func (r *request) BusinessApplication() *BusinessApplicationRequest {
	return &BusinessApplicationRequest{}
}

func (r *request) BusinessShipDetailUser() *BusinessShipDetailUserRequest {
	return &BusinessShipDetailUserRequest{}
}

func (r *request) Business() *BusinessRequest {
	return &BusinessRequest{}
}

func (r *request) BusinessPermissionAdd() *BusinessUserPermAddRequest {
	return &BusinessUserPermAddRequest{}
}

func (r *request) BusinessPermissionRemove() *BusinessUserPermRemoveRequest {
	return &BusinessUserPermRemoveRequest{}
}

func (r *request) BusinessSelect() *BusinessSelectRequest {
	return &BusinessSelectRequest{}
}

func (r *request) InviteCreate() *InviteCreateRequest {
	return &InviteCreateRequest{}
}

func (r *request) InviteDetail() *InviteDetailRequest {
	return &InviteDetailRequest{}
}

func (r *request) BusinessReject() *BusinessRejectRequest {
	return &BusinessRejectRequest{}
}
