package config

import "github.com/turistikrota/service.shared/base_roles"

type ownerRoles struct {
	Member         string
	AdminView      string
	UserAdd        string
	UserRemove     string
	UserPermAdd    string
	UserPermRemove string
	Enable         string
	Disable        string
	UserList       string
}

type roles struct {
	base_roles.Roles
	Owner ownerRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Owner: ownerRoles{
		Member:         "owner.member",
		AdminView:      "owner.admin_view",
		UserAdd:        "owner.user_add",
		UserRemove:     "owner.user_remove",
		UserPermAdd:    "owner.user_perm_add",
		UserPermRemove: "owner.user_perm_remove",
		Enable:         "owner.enable",
		Disable:        "owner.disable",
		UserList:       "owner.user_list",
	},
}
