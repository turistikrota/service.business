package service

import (
	"github.com/9ssi7/vkn"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/mixarchitecture/microp/validator"
	"github.com/ssibrahimbas/KPSPublic"
	"github.com/turistikrota/service.business/src/adapters"
	"github.com/turistikrota/service.business/src/app"
	"github.com/turistikrota/service.business/src/app/command"
	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/config"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.business/src/domain/invite"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
	I18n        *i18np.I18n
}

func NewApplication(config Config) app.Application {
	businessFactory := business.NewFactory()
	businessRepo := adapters.Mongo.NewBusiness(businessFactory, config.Mongo.GetCollection(config.App.DB.MongoBusiness.Collection))
	businessEvents := business.NewEvents(business.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
	})

	inviteFactory := invite.NewFactory()
	inviteRepo := adapters.Mongo.NewInvite(inviteFactory, config.Mongo.GetCollection(config.App.DB.MongoInvite.Collection))
	inviteEvents := invite.NewEvents(invite.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
		Urls:      config.App.Urls,
		I18n:      config.I18n,
	})

	identitySrv := KPSPublic.New()

	base := decorator.NewBase()

	vknSrv := vkn.New(vkn.Config{
		Username: config.App.Vkn.Username,
		Password: config.App.Vkn.Password,
	})

	return app.Application{
		Commands: app.Commands{
			BusinessApplication: command.NewBusinessApplicationHandler(command.BusinessApplicationHandlerConfig{
				Repo:            businessRepo,
				Factory:         businessFactory,
				IdentityService: identitySrv,
				VknService:      vknSrv,
				Events:          businessEvents,
				CqrsBase:        base,
			}),
			BusinessUserRemove: command.NewBusinessUserRemoveHandler(command.BusinessUserRemoveHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessUserPermAdd: command.NewBusinessUserPermAddHandler(command.BusinessUserPermAddHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessUserPermRemove: command.NewBusinessUserPermRemoveHandler(command.BusinessUserPermRemoveHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessEnable: command.NewBusinessEnableHandler(command.BusinessEnableConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessDisable: command.NewBusinessDisableHandler(command.BusinessDisableConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessDeleteByAdmin: command.NewAdminBusinessDeleteHandler(command.AdminBusinessDeleteConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessRecoverByAdmin: command.NewAdminBusinessRecoverHandler(command.AdminBusinessRecoverConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessVerifyByAdmin: command.NewAdminBusinessVerifyHandler(command.AdminBusinessVerifyConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			BusinessRejectByAdmin: command.NewAdminBusinessRejectHandler(command.AdminBusinessRejectConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				Events:   businessEvents,
				CqrsBase: base,
			}),
			InviteCreate: command.NewInviteCreateHandler(command.InviteCreateConfig{
				Repo:     inviteRepo,
				Factory:  inviteFactory,
				Events:   inviteEvents,
				CqrsBase: base,
			}),
			InviteUse: command.NewInviteUseHandler(command.InviteUseConfig{
				Repo:            inviteRepo,
				Factory:         inviteFactory,
				CqrsBase:        base,
				BusinessRepo:    businessRepo,
				Events:          inviteEvents,
				BusinessFactory: businessFactory,
			}),
			InviteDelete: command.NewInviteDeleteHandler(command.InviteDeleteConfig{
				Repo:     inviteRepo,
				Factory:  inviteFactory,
				Events:   inviteEvents,
				CqrsBase: base,
			}),
		},
		Queries: app.Queries{
			AdminViewBusiness: query.NewAdminViewBusinessHandler(query.AdminViewBusinessHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				CqrsBase: base,
			}),
			AdminListAll: query.NewAdminListBusinessHandler(query.AdminListBusinessHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				CqrsBase: base,
			}),
			ListMyBusinesses: query.NewListMyBusinessesHandler(query.ListMyBusinessesHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				CqrsBase: base,
			}),
			ViewBusiness: query.NewViewBusinessHandler(query.ViewBusinessHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				CqrsBase: base,
			}),
			GetWithUserBusiness: query.NewGetWithUserBusinessHandler(query.GetWithUserBusinessHandlerConfig{
				Repo:     businessRepo,
				Factory:  businessFactory,
				CqrsBase: base,
			}),
			ListMyBusinessUsers: query.NewListMyBusinessUsersQueryHandler(query.ListMyBusinessUsersQueryHandlerConfig{
				BusinessRepo:    businessRepo,
				BusinessFactory: businessFactory,
				CqrsBase:        base,
				Rpc:             config.App.Rpc,
			}),
			InviteGetByEmail: query.NewInviteGetByEmailHandler(query.InviteGetByEmailHandlerConfig{
				Repo:     inviteRepo,
				Factory:  inviteFactory,
				CqrsBase: base,
			}),
			InviteGetByUUID: query.NewInviteGetByUUIDHandler(query.InviteGetByUUIDHandlerConfig{
				Repo:     inviteRepo,
				Factory:  inviteFactory,
				CqrsBase: base,
			}),
			InviteGetByBusinessUUID: query.NewInviteGetByBusinessUUIDHandler(query.InviteGetByBusinessUUIDHandlerConfig{
				Repo:     inviteRepo,
				Factory:  inviteFactory,
				CqrsBase: base,
			}),
		},
	}
}
