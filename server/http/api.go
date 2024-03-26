package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.business/app/command"
	"github.com/turistikrota/service.business/app/query"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.business/pkg/paginate"
	"github.com/turistikrota/service.shared/server/http/auth/current_account"
	"github.com/turistikrota/service.shared/server/http/auth/current_business"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) BusinessApplication(ctx *fiber.Ctx) error {
	cmd := command.BusinessApplicationCmd{}
	h.parseBody(ctx, cmd)
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = current_user.Parse(ctx).UUID
	cmd.Locale = business.Locale(i18n.ParseLocale(ctx))
	res, err := h.app.Commands.BusinessApplication(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessSetLocale(ctx *fiber.Ctx) error {
	cmd := command.BusinessSetLocaleCmd{}
	cmd.BusinessName = current_business.Parse(ctx).NickName
	cmd.Locale = i18n.ParseLocale(ctx)
	res, err := h.app.Commands.BusinessSetLocale(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessUserRemove(ctx *fiber.Ctx) error {
	cmd := command.BusinessUserRemoveCmd{}
	h.parseParams(ctx, cmd)
	cmd.BusinessName = current_business.Parse(ctx).NickName
	cmd.AccessUserUUID = current_user.Parse(ctx).UUID
	cmd.AccessUserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.BusinessUserRemove(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AdminBusinessVerify(ctx *fiber.Ctx) error {
	cmd := command.AdminBusinessVerifyCmd{}
	h.parseParams(ctx, cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BusinessVerifyByAdmin(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AdminBusinessDelete(ctx *fiber.Ctx) error {
	cmd := command.AdminBusinessDeleteCmd{}
	h.parseParams(ctx, cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BusinessDeleteByAdmin(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AdminBusinessRecover(ctx *fiber.Ctx) error {
	cmd := command.AdminBusinessRecoverCmd{}
	h.parseParams(ctx, cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BusinessRecoverByAdmin(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AdminBusinessReject(ctx *fiber.Ctx) error {
	cmd := command.AdminBusinessRejectCmd{}
	h.parseParams(ctx, cmd)
	h.parseBody(ctx, cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BusinessRejectByAdmin(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessUserPermAdd(ctx *fiber.Ctx) error {
	cmd := command.BusinessUserPermAddCmd{}
	h.parseParams(ctx, cmd)
	h.parseBody(ctx, cmd)
	cmd.BusinessName = current_business.Parse(ctx).NickName
	cmd.AccessUserUUID = current_user.Parse(ctx).UUID
	cmd.AccessUserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.BusinessUserPermAdd(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessUserPermRemove(ctx *fiber.Ctx) error {
	cmd := command.BusinessUserPermRemoveCmd{}
	h.parseParams(ctx, cmd)
	h.parseBody(ctx, cmd)
	cmd.BusinessName = current_business.Parse(ctx).NickName
	cmd.AccessUserUUID = current_user.Parse(ctx).UUID
	cmd.AccessUserName = current_account.Parse(ctx).Name
	res, err := h.app.Commands.BusinessUserPermRemove(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) AdminViewBusiness(ctx *fiber.Ctx) error {
	query := query.AdminViewBusinessQuery{}
	h.parseParams(ctx, query)
	res, err := h.app.Queries.AdminViewBusiness(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Business)
}

func (h srv) ListMyBusinessUsers(ctx *fiber.Ctx) error {
	query := query.ListMyBusinessUsersQuery{}
	query.NickName = current_business.Parse(ctx).NickName
	query.UserName = current_account.Parse(ctx).Name
	res, err := h.app.Queries.ListMyBusinessUsers(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Users)
}

func (h srv) ListMyBusinesses(ctx *fiber.Ctx) error {
	query := query.ListMyBusinessesQuery{}
	query.UserName = current_account.Parse(ctx).Name
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.ListMyBusinesses(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Businesses)
}

func (h srv) AdminListBusinesses(ctx *fiber.Ctx) error {
	pagi := paginate.Pagination{}
	h.parseQuery(ctx, pagi)
	query := query.AdminListBusinessesQuery{
		Pagination: &pagi,
	}
	res, err := h.app.Queries.AdminListBusinesses(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.List)
}

func (h srv) ViewBusiness(ctx *fiber.Ctx) error {
	query := query.ViewBusinessQuery{}
	h.parseParams(ctx, query)
	res, err := h.app.Queries.ViewBusiness(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Business)
}

func (h srv) BusinessEnable(ctx *fiber.Ctx) error {
	cmd := command.BusinessEnableCmd{}
	cmd.BusinessName = current_business.Parse(ctx).NickName
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BusinessEnable(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessDisable(ctx *fiber.Ctx) error {
	cmd := command.BusinessDisableCmd{}
	cmd.BusinessName = current_business.Parse(ctx).NickName
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.BusinessDisable(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) BusinessSelect(ctx *fiber.Ctx) error {
	query := query.BusinessGetWithUserQuery{}
	h.parseParams(ctx, query)
	query.UserName = current_account.Parse(ctx).Name
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.BusinessGetWithUser(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	if res.Dto == nil {
		return result.ErrorDetail(Messages.Error.BusinessNotFound, map[string]interface{}{
			"mustSelect": true,
		})
	}
	ctx.Cookie(h.CreateServerSideCookie(".s.o.n", res.Dto.Entity.NickName))
	return result.Success(Messages.Success.Ok)

}

func (h srv) BusinessGetSelected(ctx *fiber.Ctx) error {
	nickName := ctx.Cookies(".s.o.n")
	if nickName == "" {
		return result.ErrorDetail(Messages.Error.BusinessNotSelected, map[string]interface{}{
			"mustSelect": true,
		})
	}
	query := query.BusinessGetWithUserQuery{}
	query.NickName = nickName
	query.UserName = current_account.Parse(ctx).Name
	query.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Queries.BusinessGetWithUser(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	if res.Dto == nil {
		return result.ErrorDetail(Messages.Error.BusinessNotFound, map[string]interface{}{
			"mustSelect": true,
		})
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Dto)
}

func (h srv) InviteCreate(ctx *fiber.Ctx) error {
	cmd := command.InviteCreateCmd{}
	h.parseBody(ctx, cmd)
	bus := current_business.Parse(ctx)
	cmd.BusinessName = bus.NickName
	cmd.BusinessUUID = bus.UUID
	cmd.CreatorUserName = current_account.Parse(ctx).Name
	cmd.CreatorUserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.InviteCreate(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) InviteDelete(ctx *fiber.Ctx) error {
	cmd := command.InviteDeleteCmd{}
	h.parseParams(ctx, cmd)
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.InviteDelete(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) InviteUse(ctx *fiber.Ctx) error {
	cmd := command.InviteUseCmd{}
	h.parseParams(ctx, cmd)
	user := current_user.Parse(ctx)
	cmd.UserName = current_account.Parse(ctx).Name
	cmd.UserUUID = user.UUID
	cmd.UserEmail = user.Email
	res, err := h.app.Commands.InviteUse(ctx.UserContext(), cmd)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res)
}

func (h srv) InviteGetByUUID(ctx *fiber.Ctx) error {
	query := query.InviteGetByUUIDQuery{}
	h.parseParams(ctx, query)
	res, err := h.app.Queries.InviteGetByUUID(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Invite)
}

func (h srv) InviteListByEmail(ctx *fiber.Ctx) error {
	query := query.InviteListByEmailQuery{}
	query.UserEmail = current_user.Parse(ctx).Email
	res, err := h.app.Queries.InviteListByEmail(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Invites)
}

func (h srv) InviteListByBusinessUUID(ctx *fiber.Ctx) error {
	query := query.InviteListByBusinessUUIDQuery{}
	query.BusinessUUID = current_business.Parse(ctx).UUID
	res, err := h.app.Queries.InviteListByBusinessUUID(ctx.UserContext(), query)
	if err != nil {
		l, a := i18n.ParseLocales(ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.Ok, res.Invites)
}
