package invite

type messages struct {
	Failed   string
	NotFound string
	Used     string
	Deleted  string
	Timeout  string
}

var I18nMessages = messages{
	Failed:   "error_invite_failed",
	NotFound: "error_invite_not_found",
	Used:     "error_invite_used",
	Deleted:  "error_invite_deleted",
	Timeout:  "error_invite_timeout",
}
