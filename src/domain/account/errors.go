package account

import "github.com/mixarchitecture/i18np"

type Errors interface {
	Failed(action string) *i18np.Error
	NotFound() *i18np.Error
}

type accountErrors struct{}

func newAccountErrors() Errors {
	return &accountErrors{}
}

func (e *accountErrors) Failed(action string) *i18np.Error {
	return i18np.NewError(I18nMessages.AccountFailed, i18np.P{"Action": action})
}

func (e *accountErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.AccountNotFound)
}
