package paginate

type Pagination struct {
	Page  *int64 `query:"page" validate:"omitempty,gt=0"`
	Limit *int64 `query:"limit" validate:"omitempty,gt=0"`
}

func (r *Pagination) Default() {
	if r.Page == nil || *r.Page <= 0 {
		r.Page = new(int64)
		*r.Page = 1
	}
	if r.Limit == nil || *r.Limit <= 0 {
		r.Limit = new(int64)
		*r.Limit = 10
	}
}
