package http

import (
	"github.com/gofiber/fiber/v2"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.business/src/config"
	"github.com/turistikrota/service.business/src/delivery/http/dto"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h Server) CurrentBusiness() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		d := dto.Request.BusinessShipDetail()
		h.parseParams(ctx, d)
		u := current_user.Parse(ctx)
		account := current_account.Parse(ctx)
		res, err := h.app.Queries.GetWithUserBusiness.Handle(ctx.UserContext(), d.ToGetWithUserBusinessQuery(u.UUID, account.Name))
		l, a := httpI18n.GetLanguagesInContext(*h.i18n, ctx)
		if err != nil {
			return result.Error(h.i18n.TranslateFromError(*err, l, a))
		}
		ctx.Locals("business", res.Business)
		return ctx.Next()
	}
}

func (h Server) parseBusiness(ctx *fiber.Ctx) *business.EntityWithUser {
	local := ctx.Locals("business")
	if local == nil {
		return nil
	}
	return local.(*business.EntityWithUser)
}

func (h Server) BusinessPermissions(perms ...string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		business := h.parseBusiness(ctx)
		if business == nil {
			return result.Error("business not found")
		}
		perms = append(perms, config.Roles.Business.Super)
		if !business.HasAnyPermissions(perms...) {
			return result.Error("access denied")
		}
		return ctx.Next()
	}
}
