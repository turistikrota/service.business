package app

import (
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	OwnerApplication        command.OwnerApplicationHandler
	OwnershipUserAdd        command.OwnershipUserAddHandler
	OwnershipUserRemove     command.OwnershipUserRemoveHandler
	OwnershipUserPermAdd    command.OwnershipUserPermAddHandler
	OwnershipUserPermRemove command.OwnershipUserPermRemoveHandler
	OwnershipEnable         command.OwnershipEnableHandler
	OwnershipDisable        command.OwnershipDisableHandler
	OwnershipVerifyByAdmin  command.AdminOwnershipVerifyHandler
	OwnershipDeleteByAdmin  command.AdminOwnershipDeleteHandler
	OwnershipRecoverByAdmin command.AdminOwnershipRecoverHandler
}

type Queries struct {
	AdminListAll         query.AdminListOwnershipHandler
	AdminViewOwnership   query.AdminViewOwnershipHandler
	ListMyOwnerships     query.ListMyOwnershipsHandler
	ListMyOwnershipUsers query.ListMyOwnershipUsersQueryHandler
	ViewOwnership        query.ViewOwnershipHandler
	GetWithUserOwnership query.GetWithUserOwnershipHandler
}
