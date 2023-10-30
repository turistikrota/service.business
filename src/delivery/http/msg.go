package http

type successMessages struct {
	OwnerApplication        string
	OwnershipUserAdd        string
	OwnershipUserRemove     string
	OwnershipUserPermAdd    string
	OwnershipUserPermRemove string
	OwnershipAdminView      string
	ListMyOwnerships        string
	AdminListAll            string
	ViewOwnership           string
	OwnershipEnable         string
	OwnershipDisable        string
	OwnershipUserList       string
	OwnershipSelect         string
	OwnershipGetSelected    string
}

type errorMessages struct {
	RequiredAuth          string
	CurrentUserAccess     string
	OwnerNotSelected      string
	AdminRoute            string
	RequiredAccountSelect string
	AccountNotFound       string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		AdminListAll:            "http_success_admin_list_all",
		OwnerApplication:        "http_success_owner_application",
		OwnershipUserAdd:        "http_success_ownership_user_add",
		OwnershipUserRemove:     "http_success_ownership_user_remove",
		OwnershipUserPermAdd:    "http_success_ownership_user_perm_add",
		OwnershipUserPermRemove: "http_success_ownership_user_perm_remove",
		OwnershipAdminView:      "http_success_ownership_admin_view",
		ListMyOwnerships:        "http_success_list_my_ownerships",
		ViewOwnership:           "http_success_view_ownership",
		OwnershipEnable:         "http_success_ownership_enable",
		OwnershipDisable:        "http_success_ownership_disable",
		OwnershipUserList:       "http_success_ownership_user_list",
		OwnershipSelect:         "http_success_ownership_select",
		OwnershipGetSelected:    "http_success_ownership_get_selected",
	},
	Error: errorMessages{
		RequiredAuth:          "http_error_required_auth",
		CurrentUserAccess:     "http_error_current_user_access",
		OwnerNotSelected:      "http_error_owner_not_selected",
		AdminRoute:            "http_error_admin_route",
		RequiredAccountSelect: "http_error_required_account_select",
		AccountNotFound:       "http_error_account_not_found",
	},
}
