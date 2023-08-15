package account

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newAccountErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}
