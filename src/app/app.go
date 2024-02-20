package app

import (
	"github.com/turistikrota/service.business/src/app/command"
	"github.com/turistikrota/service.business/src/app/query"
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
	AdminListAll            query.AdminListBusinessHandler
	AdminViewBusiness       query.AdminViewBusinessHandler
	ListMyBusinesses        query.ListMyBusinessesHandler
	ListMyBusinessUsers     query.ListMyBusinessUsersQueryHandler
	ListAsClaim             query.ListAsClaimHandler
	ViewBusiness            query.ViewBusinessHandler
	GetWithUserBusiness     query.GetWithUserBusinessHandler
	InviteGetByEmail        query.InviteGetByEmailHandler
	InviteGetByUUID         query.InviteGetByUUIDHandler
	InviteGetByBusinessUUID query.InviteGetByBusinessUUIDHandler
}
