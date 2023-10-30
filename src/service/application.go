package service

import (
	"github.com/9ssi7/vkn"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/mixarchitecture/microp/validator"
	"github.com/ssibrahimbas/KPSPublic"
	"github.com/turistikrota/service.owner/src/adapters"
	"github.com/turistikrota/service.owner/src/app"
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
}

func NewApplication(config Config) app.Application {
	ownerFactory := owner.NewFactory()
	ownerRepo := adapters.Mongo.NewOwner(ownerFactory, config.Mongo.GetCollection(config.App.DB.MongoOwner.Collection))
	ownerEvents := owner.NewEvents(owner.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
	})

	identitySrv := KPSPublic.New()

	base := decorator.NewBase()

	vknSrv := vkn.New(vkn.Config{
		Username: config.App.Vkn.Username,
		Password: config.App.Vkn.Password,
	})

	return app.Application{
		Commands: app.Commands{
			OwnerApplication: command.NewOwnerApplicationHandler(command.OwnerApplicationHandlerConfig{
				Repo:            ownerRepo,
				Factory:         ownerFactory,
				IdentityService: identitySrv,
				VknService:      vknSrv,
				Events:          ownerEvents,
				CqrsBase:        base,
			}),
			OwnershipUserAdd: command.NewOwnershipUserAddHandler(command.OwnershipUserAddHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipUserRemove: command.NewOwnershipUserRemoveHandler(command.OwnershipUserRemoveHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipUserPermAdd: command.NewOwnershipUserPermAddHandler(command.OwnershipUserPermAddHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipUserPermRemove: command.NewOwnershipUserPermRemoveHandler(command.OwnershipUserPermRemoveHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipEnable: command.NewOwnershipEnableHandler(command.OwnershipEnableConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipDisable: command.NewOwnershipDisableHandler(command.OwnershipDisableConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipDeleteByAdmin: command.NewAdminOwnershipDeleteHandler(command.AdminOwnershipDeleteConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipRecoverByAdmin: command.NewAdminOwnershipRecoverHandler(command.AdminOwnershipRecoverConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
			OwnershipVerifyByAdmin: command.NewAdminOwnershipVerifyHandler(command.AdminOwnershipVerifyConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				Events:   ownerEvents,
				CqrsBase: base,
			}),
		},
		Queries: app.Queries{
			AdminViewOwnership: query.NewAdminViewOwnershipHandler(query.AdminViewOwnershipHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				CqrsBase: base,
			}),
			AdminListAll: query.NewAdminListOwnershipHandler(query.AdminListOwnershipHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				CqrsBase: base,
			}),
			ListMyOwnerships: query.NewListMyOwnershipsHandler(query.ListMyOwnershipsHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				CqrsBase: base,
			}),
			ViewOwnership: query.NewViewOwnershipHandler(query.ViewOwnershipHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				CqrsBase: base,
			}),
			GetWithUserOwnership: query.NewGetWithUserOwnershipHandler(query.GetWithUserOwnershipHandlerConfig{
				Repo:     ownerRepo,
				Factory:  ownerFactory,
				CqrsBase: base,
			}),
			ListMyOwnershipUsers: query.NewListMyOwnershipUsersQueryHandler(query.ListMyOwnershipUsersQueryHandlerConfig{
				OwnerRepo:    ownerRepo,
				OwnerFactory: ownerFactory,
				CqrsBase:     base,
			}),
		},
	}
}
