package service

import (
	"github.com/9ssi7/vkn"
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/ssibrahimbas/KPSPublic"
	"github.com/turistikrota/service.business/app"
	"github.com/turistikrota/service.business/app/command"
	"github.com/turistikrota/service.business/app/query"
	"github.com/turistikrota/service.business/config"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.business/domains/invite"
	"github.com/turistikrota/service.shared/cipher"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validation.Validator
	I18n        *i18np.I18n
}

func NewApplication(config Config) app.Application {
	businessFactory := business.NewFactory()
	businessRepo := business.NewRepository(businessFactory, config.Mongo.GetCollection(config.App.DB.MongoBusiness.Collection))
	businessEvents := business.NewEvents(business.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
		I18n:      config.I18n,
	})

	inviteFactory := invite.NewFactory()
	inviteRepo := invite.NewRepository(inviteFactory, config.Mongo.GetCollection(config.App.DB.MongoInvite.Collection))
	inviteEvents := invite.NewEvents(invite.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
		Urls:      config.App.Urls,
		I18n:      config.I18n,
	})

	identitySrv := KPSPublic.New()

	taxIdentitySrv := vkn.New(vkn.Config{
		Username: config.App.Vkn.Username,
		Password: config.App.Vkn.Password,
	})
	cipher := cipher.New(config.App.Cipher.Key, config.App.Cipher.IV)
	return app.Application{
		Commands: app.Commands{
			BusinessApplication:    command.NewBusinessApplicationHandler(businessRepo, businessFactory, businessEvents, identitySrv, taxIdentitySrv, cipher),
			BusinessUserRemove:     command.NewBusinessUserRemoveHandler(businessRepo, businessFactory, businessEvents),
			BusinessUserPermAdd:    command.NewBusinessUserPermAddHandler(businessRepo, businessFactory, businessEvents),
			BusinessUserPermRemove: command.NewBusinessUserPermRemoveHandler(businessRepo, businessFactory, businessEvents),
			BusinessEnable:         command.NewBusinessEnableHandler(businessRepo, businessFactory, businessEvents),
			BusinessDisable:        command.NewBusinessDisableHandler(businessRepo, businessFactory, businessEvents),
			BusinessVerifyByAdmin:  command.NewAdminBusinessVerifyHandler(businessRepo, businessFactory, businessEvents),
			BusinessDeleteByAdmin:  command.NewAdminBusinessDeleteHandler(businessRepo, businessFactory, businessEvents),
			BusinessSetLocale:      command.NewBusinessSetLocaleHandler(businessRepo, businessFactory),
			BusinessRecoverByAdmin: command.NewAdminBusinessRecoverHandler(businessRepo, businessFactory, businessEvents),
			BusinessRejectByAdmin:  command.NewAdminBusinessRejectHandler(businessRepo, businessFactory, businessEvents),
			InviteCreate:           command.NewInviteCreateHandler(inviteRepo, inviteFactory, inviteEvents),
			InviteUse:              command.NewInviteUseHandler(inviteRepo, inviteFactory, inviteEvents, businessRepo, businessFactory),
			InviteDelete:           command.NewInviteDeleteHandler(inviteRepo, inviteEvents),
		},
		Queries: app.Queries{
			AdminListBusinesses:      query.NewAdminListBusinessesHandler(businessRepo),
			AdminViewBusiness:        query.NewAdminViewBusinessHandler(businessRepo),
			ListMyBusinesses:         query.NewListMyBusinessesHandler(businessRepo, businessFactory),
			ListMyBusinessUsers:      query.NewListMyBusinessUsersHandler(businessRepo, businessFactory, config.App.Rpc),
			ListAsClaim:              query.NewListAsClaimHandler(businessRepo, businessFactory),
			ViewBusiness:             query.NewViewBusinessHandler(businessRepo),
			ViewMyBusiness:           query.NewViewMyBusinessHandler(businessRepo),
			BusinessGetWithUser:      query.NewBusinessGetWithUserHandler(businessRepo),
			InviteListByEmail:        query.NewInviteListByEmailHandler(inviteRepo, inviteFactory),
			InviteGetByUUID:          query.NewInviteGetByUUIDHandler(inviteRepo, inviteFactory),
			InviteListByBusinessUUID: query.NewInviteListByBusinessUUIDHandler(inviteRepo, inviteFactory),
		},
	}
}
