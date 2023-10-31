package invite

import "github.com/mixarchitecture/i18np"

type Errors interface {
	InvalidUUID() *i18np.Error
	Failed(operation string) *i18np.Error
	NotFound() *i18np.Error
	Used() *i18np.Error
	Deleted() *i18np.Error
	Timeout() *i18np.Error
}

type inviteErrors struct{}

func newInviteErrors() Errors {
	return &inviteErrors{}
}

func (e *inviteErrors) InvalidUUID() *i18np.Error {
	return i18np.NewError(I18nMessages.InvalidUUID)
}

func (e *inviteErrors) Failed(operation string) *i18np.Error {
	return i18np.NewError(I18nMessages.Failed, i18np.P{"Operation": operation})
}

func (e *inviteErrors) NotFound() *i18np.Error {
	return i18np.NewError(I18nMessages.NotFound)
}

func (e *inviteErrors) Used() *i18np.Error {
	return i18np.NewError(I18nMessages.Used)
}

func (e *inviteErrors) Deleted() *i18np.Error {
	return i18np.NewError(I18nMessages.Deleted)
}

func (e *inviteErrors) Timeout() *i18np.Error {
	return i18np.NewError(I18nMessages.Timeout)
}
