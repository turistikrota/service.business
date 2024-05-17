package invite

type messages struct {
	InvalidUUID string
	Failed      string
	NotFound    string
	Used        string
	Deleted     string
	Timeout     string

	InviteSubject string
	EmailMismatch string
}

var Messages = messages{
	InvalidUUID: "error_invite_invalid_uuid",
	Failed:      "error_invite_failed",
	NotFound:    "error_invite_not_found",
	Used:        "error_invite_used",
	Deleted:     "error_invite_deleted",
	Timeout:     "error_invite_timeout",

	InviteSubject: "invite_subject",
	EmailMismatch: "error_invite_email_mismatch",
}
