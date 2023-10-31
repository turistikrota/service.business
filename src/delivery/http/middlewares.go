package http

import (
	"github.com/gofiber/fiber/v2"
	httpI18n "github.com/mixarchitecture/microp/server/http/i18n"
	"github.com/mixarchitecture/microp/server/http/result"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.owner/src/delivery/http/dto"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h Server) CurrentOwner() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		d := dto.Request.OwnerShipDetail()
		h.parseParams(ctx, d)
		u := current_user.Parse(ctx)
		account := current_account.Parse(ctx)
		res, err := h.app.Queries.GetWithUserOwnership.Handle(ctx.UserContext(), d.ToGetWithUserOwnershipQuery(u.UUID, account.Name))
		l, a := httpI18n.GetLanguagesInContext(*h.i18n, ctx)
		if err != nil {
			return result.Error(h.i18n.TranslateFromError(*err, l, a))
		}
		ctx.Locals("owner", res.Ownership)
		return ctx.Next()
	}
}

func (h Server) parseOwner(ctx *fiber.Ctx) *owner.EntityWithUser {
	local := ctx.Locals("owner")
	if local == nil {
		return nil
	}
	return local.(*owner.EntityWithUser)
}

func (h Server) OwnerPermissions(perms ...string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ownership := h.parseOwner(ctx)
		if ownership == nil {
			return result.Error("ownership not found")
		}
		perms = append(perms, config.Roles.Owner.Super)
		if !ownership.HasAnyPermissions(perms...) {
			return result.Error("access denied")
		}
		return ctx.Next()
	}
}
