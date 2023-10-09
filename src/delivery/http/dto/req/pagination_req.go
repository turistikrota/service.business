package req

type PaginationRequest struct {
	Page  *int64 `query:"page" validate:"omitempty,gt=0"`
	Limit *int64 `query:"limit" validate:"omitempty,gt=0"`
}

func (r *PaginationRequest) Default() {
	if r.Page == nil {
		r.Page = new(int64)
		*r.Page = 1
	}
	if r.Limit == nil {
		r.Limit = new(int64)
		*r.Limit = 10
	}
}

func (r *request) Pagination() *PaginationRequest {
	return &PaginationRequest{}
}
