package app

import (
	"github.com/turistikrota/service.business/app/command"
	"github.com/turistikrota/service.business/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	BusinessApplication    command.BusinessApplicationHandler
	BusinessUserRemove     command.BusinessUserRemoveHandler
	BusinessUserPermAdd    command.BusinessUserPermAddHandler
	BusinessUserPermRemove command.BusinessUserPermRemoveHandler
	BusinessEnable         command.BusinessEnableHandler
	BusinessDisable        command.BusinessDisableHandler
	BusinessVerifyByAdmin  command.AdminBusinessVerifyHandler
	BusinessDeleteByAdmin  command.AdminBusinessDeleteHandler
	BusinessSetLocale      command.BusinessSetLocaleHandler
	BusinessRecoverByAdmin command.AdminBusinessRecoverHandler
	BusinessRejectByAdmin  command.AdminBusinessRejectHandler
	InviteCreate           command.InviteCreateHandler
	InviteUse              command.InviteUseHandler
	InviteDelete           command.InviteDeleteHandler
}

type Queries struct {
	AdminListBusinesses      query.AdminListBusinessesHandler
	AdminViewBusiness        query.AdminViewBusinessHandler
	ListMyBusinesses         query.ListMyBusinessesHandler
	ListMyBusinessUsers      query.ListMyBusinessUsersHandler
	ListAsClaim              query.ListAsClaimHandler
	ViewBusiness             query.ViewBusinessHandler
	BusinessGetWithUser      query.BusinessGetWithUserHandler
	InviteListByEmail        query.InviteListByEmailHandler
	InviteGetByUUID          query.InviteGetByUUIDHandler
	InviteListByBusinessUUID query.InviteListByBusinessUUIDHandler
}
