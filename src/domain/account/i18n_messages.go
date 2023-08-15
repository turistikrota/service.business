package account

type messages struct {
	AccountFailed   string
	AccountNotFound string
}

var I18nMessages = messages{
	AccountFailed:   "error_account_failed",
	AccountNotFound: "error_account_not_found",
}
