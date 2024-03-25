package config

import "github.com/turistikrota/service.shared/base_roles"

type businessRoles struct {
	Member         string
	AdminView      string
	AdminList      string
	AdminReject    string
	AdminVerify    string
	AdminDelete    string
	AdminRecover   string
	UserAdd        string
	UserRemove     string
	UserPermAdd    string
	UserPermRemove string
	Enable         string
	Disable        string
	UserList       string
	LocaleSet      string
	Super          string
	InviteCreate   string
	InviteDelete   string
	InviteView     string
	UploadAvatar   string
	UploadCover    string
}

type roles struct {
	base_roles.Roles
	Business businessRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Business: businessRoles{
		Super:          "business.super",
		Member:         "business.member",
		AdminView:      "business.admin_view",
		AdminList:      "business.admin_list",
		AdminReject:    "business.admin_reject",
		AdminVerify:    "business.admin_verify",
		AdminDelete:    "business.admin_delete",
		AdminRecover:   "business.admin_recover",
		UserAdd:        "business.user_add",
		UserRemove:     "business.user_remove",
		LocaleSet:      "business.locale_set",
		UserPermAdd:    "business.user_perm_add",
		UserPermRemove: "business.user_perm_remove",
		Enable:         "business.enable",
		Disable:        "business.disable",
		UserList:       "business.user_list",
		InviteCreate:   "business.invite_create",
		InviteDelete:   "business.invite_delete",
		InviteView:     "business.invite_view",
		UploadAvatar:   "business.upload.avatar",
		UploadCover:    "business.upload.cover",
	},
}
