package http

type successMessages struct {
	BusinessApplication    string
	BusinessUserAdd        string
	BusinessUserRemove     string
	BusinessUserPermAdd    string
	BusinessUserPermRemove string
	BusinessAdminView      string
	ListMyBusinesses       string
	AdminListAll           string
	ViewBusiness           string
	BusinessEnable         string
	BusinessDisable        string
	BusinessUserList       string
	BusinessSelect         string
	BusinessGetSelected    string
	Ok                     string
	BusinessSetLocale      string
}

type errorMessages struct {
	RequiredAuth          string
	CurrentUserAccess     string
	BusinessNotSelected   string
	AdminRoute            string
	RequiredAccountSelect string
	AccountNotFound       string
	BusinessNotFound      string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		AdminListAll:           "http_success_admin_list_all",
		BusinessApplication:    "http_success_business_application",
		BusinessUserAdd:        "http_success_business_user_add",
		BusinessUserRemove:     "http_success_business_user_remove",
		BusinessUserPermAdd:    "http_success_business_user_perm_add",
		BusinessUserPermRemove: "http_success_business_user_perm_remove",
		BusinessAdminView:      "http_success_business_admin_view",
		ListMyBusinesses:       "http_success_list_my_businesss",
		ViewBusiness:           "http_success_view_business",
		BusinessEnable:         "http_success_business_enable",
		BusinessDisable:        "http_success_business_disable",
		BusinessUserList:       "http_success_business_user_list",
		BusinessSelect:         "http_success_business_select",
		BusinessGetSelected:    "http_success_business_get_selected",
		Ok:                     "http_success_ok",
		BusinessSetLocale:      "http_success_business_set_locale",
	},
	Error: errorMessages{
		RequiredAuth:          "http_error_required_auth",
		CurrentUserAccess:     "http_error_current_user_access",
		BusinessNotSelected:   "http_error_business_not_selected",
		AdminRoute:            "http_error_admin_route",
		RequiredAccountSelect: "http_error_required_account_select",
		AccountNotFound:       "http_error_account_not_found",
		BusinessNotFound:      "http_error_business_not_found",
	},
}
