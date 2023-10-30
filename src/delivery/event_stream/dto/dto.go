package dto

type Dto interface {
}

type dto struct{}

func New() Dto {
	return &dto{}
}
